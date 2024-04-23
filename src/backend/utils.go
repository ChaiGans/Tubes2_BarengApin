package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// fetchLinks uses Colly to scrape Wikipedia for links on a given page.
func fetchLinks(startUrl string) ([]string, error) {
    links := []string{}

    // Define the namespaces to be excluded
    excludedNamespaces := []string{
        "Category:", "Wikipedia:", "File:", "Help:", "Portal:",
        "Special:", "Talk:", "User:","Template:", "Template_talk:", "Main_Page",
    }

    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
        colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"),
    )

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") {
            exclude := false
            for _, namespace := range excludedNamespaces {
                if strings.Contains(link, namespace) {
                    exclude = true
                    break
                }
            }
            if (!exclude) {
                fullLink := "https://en.wikipedia.org" + link
                if (!slices.Contains(links, fullLink)) {
                    links = append(links, fullLink)
                }
            }
        }
    })

    // fmt.Println("Visiting URL:", startUrl)
    err := c.Visit(startUrl)
    if err != nil {
        return nil, err
    }
    c.Wait()
    return links, nil
}


// INI VERSI 2

// List of user agents to rotate 
var userAgents = []string{
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
    "Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
    "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",
    "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
}

func getRandomUserAgent() string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return userAgents[r.Intn(len(userAgents))]
}

func scrapeWikipediaLinksAsync(url string) ([]string, error) {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
        colly.Async(true), 
    )

    c.Limit(&colly.LimitRule{
        DomainGlob: "*wikipedia.org",
        Parallelism: 1,
    })

    c.UserAgent = getRandomUserAgent()

    var links []string 
    var mu sync.Mutex 

    excludedNamespaces := []string{
        "Category:", "Wikipedia:", "File:", "Help:", "Portal:",
        "Special:", "Talk:", "User:", "Template:", "Template_talk:", "Main_Page",
    }

    wikiArticleRegex := regexp.MustCompile(`^https?://en\.wikipedia\.org/wiki/([^#:\s]+)$`)

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("href"))
        if wikiArticleRegex.MatchString(link) {
            exclude := false
            for _, namespace := range excludedNamespaces {
                if strings.Contains(link, namespace) {
                    exclude = true
                    break
                }
            }
            if !exclude {
                mu.Lock()
                links = append(links, link) 
                mu.Unlock()
            }
        }
    })

    var err error
    for i := 0; i < 3; i++ {
        err = c.Visit(url)
        c.OnError(func(r *colly.Response, err error) {
            fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
        })
        if err == nil {
            // fmt.Println("done")
            break 
        }
        fmt.Println("Retrying:", url, "Attempt:", i+1)
        time.Sleep(time.Second * 1) 
    }

    c.Wait() 
    return links, err 
}

package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

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

func scrapeWikipediaLinks(url string) ([]string, error) {
    c := colly.NewCollector(
        colly.AllowedDomains("en.wikipedia.org"),
        colly.Async(true), 
    )

    c.Limit(&colly.LimitRule{
        DomainGlob: "*wikipedia.org",
        Parallelism: 1,
    })

    c.UserAgent = getRandomUserAgent()
    c.CacheDir = "./cache"
    
    var links []string 
    var mu sync.Mutex 

    excludedNamespaces := []string{
        "Category:", "Wikipedia:", "File:", "Help:", "Portal:",
        "Special:", "Talk:", "User:", "Template:", "Template_talk:", "Main_Page",
    }

    wikiArticleRegex := regexp.MustCompile(`^https?://en\.wikipedia\.org/wiki/([^#:\s]+)$`)

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        if !checkDisplayNone(e) {
            return 
        }

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
        if err == nil {
            break
        }
        c.OnError(func(r *colly.Response, err error) {
            fmt.Println("Request URL:", r.Request.URL, "Error:", err)
        })
        sleepDuration := time.Millisecond * time.Duration(500 * (1 << i)) 
        fmt.Println("Retrying:", url, "Attempt:", i+1, "after", sleepDuration)
        time.Sleep(sleepDuration)
    }

    c.Wait() 
    return MakeUnique(links), err 
}
func contains(slice []string, item string) bool {
    for _, s := range slice {
        if s == item {
            return true
        }
    }
    return false
}

func checkDisplayNone(e *colly.HTMLElement) bool {
    class := e.Attr("class")
    classes := strings.Split(class, " ")
    if contains(classes, "nowraplinks") {
        return false
    }

    // Check parent elements 
    for parent := e.DOM.Parent(); parent.Length() != 0; parent = parent.Parent() {
        parentClass, found := parent.Attr("class")
        parentClass = strings.ReplaceAll(parentClass, " ", "")
        if found && strings.Contains(parentClass, "nowraplinks") {
            // fmt.Println(e.Attr("href"))
            return false
        }
    }
    return true
}


func MakeUnique(strings []string) []string {
    uniqueMap := make(map[string]bool)
    var uniqueStrings []string     

    for _, str := range strings {
        if _, exists := uniqueMap[str]; !exists {
            uniqueMap[str] = true
            uniqueStrings = append(uniqueStrings, str)
        }
    }

    return uniqueStrings
}
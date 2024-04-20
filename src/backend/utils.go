package main

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

// fetchLinks uses Colly to scrape Wikipedia for links on a given page.
func fetchLinks(startUrl string) ([]string, error) {
    links := make([]string, 0)

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
                links = append(links, fullLink)
            }
        }
    })

    // fmt.Println("Visiting URL:", startUrl) // Log the URL being visited.
    err := c.Visit(startUrl)
    if err != nil {
        return nil, err
    }
    c.Wait()
    return links, nil
}

// func scrapeWikipediaLinksAsync(url string) ([]string, error) {
//     // Initialize the Colly collector
//     c := colly.NewCollector(
//         colly.AllowedDomains("en.wikipedia.org","www.wikipedia.org"),
//         colly.Async(true), // Enable asynchronous scraping
//     )

//     // Configure the rate limit to be gentle with Wikipedia's servers
//     c.Limit(&colly.LimitRule{
//         DomainGlob:  "*wikipedia.org",
//         Parallelism: 5, // Limit the number of parallel requests
//     })

//     // Set a consistent User-Agent for all requests
//     c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

//     // Cache responses to avoid re-fetching unchanged pages
//     c.CacheDir = "./cache"

//     var links []string // List to store the found links
//     var mu sync.Mutex // Mutex to handle concurrent link append operations safely

//     // Define the namespaces to be excluded
//     excludedNamespaces := []string{
//         "Category:", "Wikipedia:", "File:", "Help:", "Portal:",
//         "Special:", "Talk:", "User:","Template:", "Template_talk:", "Main_Page",
//     }

//     // Regex to filter only valid article links, excluding the namespaces
//     wikiArticleRegex := regexp.MustCompile(`^https?://en\.wikipedia\.org/wiki/([^#:\s]+)$`)

//     // Handle each anchor tag found during scraping
//     c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//         link := e.Request.AbsoluteURL(e.Attr("href"))
//         if wikiArticleRegex.MatchString(link) {
//             exclude := false
//             for _, namespace := range excludedNamespaces {
//                 if strings.Contains(link, namespace) {
//                     exclude = true
//                     break
//                 }
//             }
//             if !exclude {
//                 mu.Lock()
//                 links = append(links, link) // Safely append link to the slice
//                 mu.Unlock()
//             }
//         }
//     })

//     // Error handling
//     c.OnError(func(r *colly.Response, err error) {
//         fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "Error:", err)
//     })

//     // Visit the URL
//     err := c.Visit(url)
//     c.Wait() // Wait for all asynchronous tasks to complete
//     return links, err // Return the links and any error encountered
// }
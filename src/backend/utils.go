package main

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

// fetchLinks uses Colly to scrape Wikipedia for links on a given page.
func fetchLinks(articleTitle string) ([]string, error) {
    links := make([]string, 0)
    c := colly.NewCollector(
        colly.AllowedDomains("id.wikipedia.org"),
    )

    // Look for all <a> tags that link to other Wikipedia articles
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        // Check if the link is a valid Wikipedia article link
        if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") {
            fullLink := "https://id.wikipedia.org" + link
            links = append(links, fullLink)
        }
    })

    // Start scraping on the article page
    startUrl := "https://id.wikipedia.org/wiki/" + strings.ReplaceAll(articleTitle, " ", "_")
    err := c.Visit(startUrl)
    if err != nil {
        return nil, err
    }

    c.Wait() // Wait for all requests to finish
    return links, nil
}

package main

import (
	"fmt"
	"sync"
	"time"
)


func bfs(startURL, targetURL string) ([]string, int, int, int64, error) {
    startTime := time.Now()
    queue := [][]string{{startURL}}
    visited := make(map[string]bool)
    visited[startURL] = true
    numChecked := 0

    var mu sync.Mutex
    var wg sync.WaitGroup
    results := make(chan []string)

    processLinks := func(path []string) {
        defer wg.Done()
        currentURL := path[len(path)-1]
        numChecked++

        links, err := scrapeWikipediaLinksAsync(currentURL)
        if err != nil {
            fmt.Println("Error scraping:", err)
            return
        }

        mu.Lock()
        for _, link := range links {
            if link == targetURL {
                results <- append(path, link)
                mu.Unlock()
                return
            }
            if !visited[link] {
                visited[link] = true
                newPath := append([]string{}, path...)
                newPath = append(newPath, link)
                queue = append(queue, newPath)
            }
        }
        mu.Unlock()
    }

    for len(queue) > 0 {
        nextQueue := [][]string{}
        for _, path := range queue {
            wg.Add(1)
            go processLinks(path)
        }
        wg.Wait() // Wait for all goroutines to finish before moving to the next level
        queue = nextQueue
    }

    select {
    case path := <-results:
        duration := time.Since(startTime).Milliseconds()
        return path, len(path), numChecked, duration, nil
    default:
        return nil, 0, numChecked, time.Since(startTime).Milliseconds(), fmt.Errorf("no path found from %s to %s", startURL, targetURL)
    }
}


func main() {
    startURL := "https://en.wikipedia.org/wiki/Indonesia"
    targetURL := "https://en.wikipedia.org/wiki/Pornography"
    fmt.Println("backend running")
    // links, err := scrapeWikipediaLinksAsync(startURL)
    // fmt.Println(len(links))
    path, length, numChecked, duration, err := bfs(startURL, targetURL)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Path found:")
    for _, url := range path {
        fmt.Println(url)
    }
    fmt.Printf("Path length: %d, Articles checked: %d, Time taken: %d ms\n", length, numChecked, duration)
}
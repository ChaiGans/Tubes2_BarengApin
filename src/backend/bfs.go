package main

import (
	"fmt"
	"sync"
	"time"
)

func bfs(startURL, targetURL string) ([]string, int, int, int64, error) {
    startTime := time.Now()
    var wg sync.WaitGroup
    var mu sync.Mutex
    visited := make(map[string]bool)
    visited[startURL] = true
    numChecked := 0

    queue := [][]string{{startURL}}
    results := make(chan []string)
    done := make(chan bool)

    // Worker function 
    processURLs := func(paths [][]string) {
        defer wg.Done()
        localQueue := [][]string{}

        for _, path := range paths {
            currentURL := path[len(path)-1]

            mu.Lock()
            numChecked++
            mu.Unlock()

            links, err := fetchLinks(currentURL)
            if err != nil {
                fmt.Println("Error scraping:", err)
                continue
            }

            for _, link := range links {
                if link == targetURL {
                    results <- append(path, link)
                    return
                }

                mu.Lock()
                if !visited[link] {
                    visited[link] = true
                    newPath := append([]string{}, path...)
                    newPath = append(newPath, link)
                    localQueue = append(localQueue, newPath)
                }
                mu.Unlock()
            }
        }

        mu.Lock()
        queue = append(queue, localQueue...)
        mu.Unlock()
    }

    go func() {
        for len(queue) > 0 {
            currentBatch := queue
            queue = nil // Clear Queue

            wg.Add(1)
            go processURLs(currentBatch) // Process current batch to new goroutine

            wg.Wait() // wait all goroutines  finish before processing the next level
        }
        close(done)
    }()

    select {
    case path := <-results:
        duration := time.Since(startTime).Milliseconds()
        return path, len(path), numChecked, duration, nil
    case <-done:
        return nil, 0, numChecked, time.Since(startTime).Milliseconds(), fmt.Errorf("no path found from %s to %s", startURL, targetURL)
    }
}
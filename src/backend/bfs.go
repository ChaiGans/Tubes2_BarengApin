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

    // Concurrency limiter
    maxGoroutines := 10
    semaphore := make(chan struct{}, maxGoroutines)

    // Worker function 
    processURLs := func(paths [][]string) {
        defer wg.Done()
        defer func() { <-semaphore }() 
        localQueue := [][]string{}

        for _, path := range paths {
            currentURL := path[len(path)-1]

            links, err := scrapeWikipediaLinks(currentURL)

            // fmt.Println("Done Scrapping", len(links), currentURL)
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
                    numChecked++
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
            queue = nil 

            for _, paths := range currentBatch {
                semaphore <- struct{}{} 
                wg.Add(1)
                go processURLs([][]string{paths})
            }

            wg.Wait() 
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



func bfsMultiPath(startURL, targetURL string) ([][]string, int, int, int64, error) {
    startTime := time.Now()
    var wg sync.WaitGroup
    var mu sync.Mutex
    visited := make(map[string]bool)
    visited[startURL] = true
    numChecked := 0

    queue := [][]string{{startURL}}
    resultsMap := make(map[string][]string) 
    var results [][]string
    foundLevel := -1

    maxGoroutines := 20
    goroutineSem := make(chan struct{}, maxGoroutines)

    for len(queue) > 0 {
        currentLevel := queue
        queue = nil
        var levelQueue [][]string

        for _, path := range currentLevel {
            wg.Add(1)
            goroutineSem <- struct{}{}

            go func(path []string) {
                defer wg.Done()
                defer func() { <-goroutineSem }()

                currentURL := path[len(path)-1]
                if foundLevel != -1 && len(path) > foundLevel {
                    return
                }

                links, err := scrapeWikipediaLinks(currentURL)
                if err != nil {
                    fmt.Println("Error scraping:", err)
                    return
                }

                mu.Lock()
                defer mu.Unlock()

                for _, link := range links {
                    if link == targetURL && (foundLevel == -1 || len(path) == foundLevel) {
                        pathSignature := createPathSignature(append(path, link))
                        if _, exists := resultsMap[pathSignature]; !exists {
                            newPath := append([]string(nil), append(path, link)...)
                            resultsMap[pathSignature] = newPath
                            results = append(results, newPath)
                        }
                        foundLevel = len(path)
                    }
                    if foundLevel != -1 && len(path) >= foundLevel {
                        if !visited[link] {
                            visited[link] = true
                            numChecked++
                        }
                        continue
                    }
                    if !visited[link] {
                        visited[link] = true
                        levelQueue = append(levelQueue, append([]string(nil), append(path, link)...))
                        numChecked++
                    }
                }
            }(path)
        }

        wg.Wait()
        queue = levelQueue
    }

    duration := time.Since(startTime).Milliseconds()
    if len(results) > 0 {
        pathLength := len(results[0])
        return results, pathLength, numChecked, duration, nil
    }
    return nil, 0, numChecked, duration, fmt.Errorf("no path found from %s to %s", startURL, targetURL)
}

func createPathSignature(path []string) string {
    signature := ""
    for _, p := range path {
        signature += p + "|"
    }
    return signature
}


// func main() {
//     runtime.GOMAXPROCS(runtime.NumCPU())
//     startURL := "https://en.wikipedia.org/wiki/Joko_Widodo"
//     targetURL := "https://en.wikipedia.org/wiki/Korea"
//     fmt.Println("backend running")
//     // links, err := scrapeWikipediaLinksAsync(startURL)
//     // for _, url := range links {
//     //     fmt.Println(url)
//     // }
//     // fmt.Println(len(links))
//     // path, length, numChecked,numRequested, duration, err := bfs(startURL, targetURL)
//     path, length, numChecked, duration, err := bfsMultiPath(startURL, targetURL)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }

//     fmt.Println("Paths found:")
//     for i, path := range path {
//         fmt.Printf("Path %d:\n", i+1)
//         for _, url := range path {
//             fmt.Println(url)
//         }
//     }
//     fmt.Printf("Number of paths: %d\n", len(path))
//     fmt.Printf("Path length: %d, Articles checked: %d, Time taken: %d ms\n", length, numChecked, duration)
// }
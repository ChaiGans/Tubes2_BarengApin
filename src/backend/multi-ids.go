package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Depth-Limit Search
func MultiDLS(depth int, startLink, goalLink string, currentDepth int, path []string, visited map[string]bool, cache *LocalCache, multiple_path_save *[][]string) (error) {
    if currentDepth > depth {
        return nil
    }

    visited[startLink] = true
    // log.Printf("Visiting: %s", startLink)

    links, ok := cache.Get(startLink)
    if !ok {
        var err error
        links, err = scrapeWikipediaLinks(startLink)
        if err != nil {
            return fmt.Errorf("error fetching links at title %s: %v", startLink, err)
        }
        cache.Set(startLink, links)
    }

    for _, link := range links {
        if !visited[link] {
            // log.Printf("Visiting new link: %s, Total links visited: %d, At depth : %d, Current len(multiplepathsave) : %d", link, len(visited), depth, len(*multiple_path_save))
            visited[link] = true

            if strings.EqualFold(link, goalLink) {
                *multiple_path_save = append(*multiple_path_save, append(path, link))
            } else {
                new_path := append([]string{}, path...)
                new_path = append(new_path, link)
                err := MultiDLS(depth, link, goalLink, currentDepth+1, new_path, visited, cache, multiple_path_save)
                if err != nil {
                    log.Printf("Error at link %s", link)
                }
		    }
        } 
        // else {
        //     // log.Printf("Skipping current link already visited: %s, Total links visited: %d", link, len(visited))
        // }

        visited[link] = false
    }

    if len(*multiple_path_save) > 0 {
        return nil
    }
    return fmt.Errorf("goal not found at depth %d", depth)
}

func MultiIDS(startLink, goalLink string) ([][]string, error, int, int64) {
    var i int = 0
	multiple_path_save := [][]string{}
    cache := NewLocalCache()
    
    for {
        visited := make(map[string]bool)
        start := time.Now()
        err := MultiDLS(i, startLink, goalLink, 0, []string{startLink}, visited, cache, &multiple_path_save)
        elapsed := time.Since(start).Milliseconds()
        if err == nil {
            return multiple_path_save, nil, len(visited), elapsed
        }
        // log.Printf("No result at depth %d, error: %v", i, err)
        i++
        if i > 20 || len(multiple_path_save) > 0 {
            break
        }
    }
    return nil, fmt.Errorf("goal not found after depth %d", i), 0, 0
}
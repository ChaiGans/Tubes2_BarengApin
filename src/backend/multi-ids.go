package main

import (
	"fmt"
	"strings"
	"time"
)

// Depth-Limit Search
func MultiDLS(depth int, startLink, goalLink string, currentDepth int, path []string, visited map[string]bool, totalLinkIterate *int, cache *LocalCache, multiple_path_save *[][]string) (error) {
    if currentDepth > depth {
        return nil
    }

    visited[startLink] = true
    // log.Printf("Visiting: %s", startLink)

    links, ok := cache.Get(startLink)
    if !ok {
        var err error
        links, err = fetchLinks(startLink)
        if err != nil {
            return fmt.Errorf("error fetching links at title %s: %v", startLink, err)
        }
        cache.Set(startLink, links)
    }

    for _, link := range links {
        if !visited[link] {
            *totalLinkIterate++
            // log.Printf("Visiting new link: %s, Total links visited: %d, At depth : %d, Current found: %d, Current path length : %d", link, *totalLinkIterate, depth, len(*multiple_path_save), len(path))
            visited[link] = true
        }

        if strings.EqualFold(link, goalLink) {
			*multiple_path_save = append(*multiple_path_save, append(path, link))
        } else {
			new_path := append([]string{}, path...)
        	new_path = append(new_path, link)
			err := MultiDLS(depth, link, goalLink, currentDepth+1, new_path, visited, totalLinkIterate, cache, multiple_path_save)
            if err != nil {
                // log.Printf("Error at link %s: %v", link, err)
            }
		}
    }

    if len(*multiple_path_save) > 0 {
        return nil
    }
    return fmt.Errorf("goal not found at depth %d", depth)
}

func MultiIDS(startLink, goalLink string) ([][]string, error, int, int64) {
    var i int = 0
    var iterationNumber int = 0
	multiple_path_save := [][]string{}
    visited := make(map[string]bool)
    cache := NewLocalCache()

    for {
        start := time.Now()
        err := MultiDLS(i, startLink, goalLink, 0, []string{startLink}, visited, &iterationNumber, cache, &multiple_path_save)
        elapsed := time.Since(start).Milliseconds()
        if err == nil {
            return multiple_path_save, nil, iterationNumber, elapsed
        }
        // log.Printf("No result at depth %d, error: %v", i, err)
        i++
        if i > 20 || len(multiple_path_save) > 0 {
            break
        }
    }
    return nil, fmt.Errorf("goal not found after depth %d", i), iterationNumber, 0
}

func makeUnique(arrays [][]string) [][]string {
	uniqueMap := make(map[string]bool)
	var uniqueArrays [][]string

	for _, arr := range arrays {
		key := fmt.Sprintf("%v", arr)
		if _, ok := uniqueMap[key]; !ok {
			uniqueMap[key] = true
			uniqueArrays = append(uniqueArrays, arr)
		}
	}

	return uniqueArrays
}
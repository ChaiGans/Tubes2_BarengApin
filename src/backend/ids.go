package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// Define a struct to hold the cache and its lock
type LocalCache struct {
    mu    sync.RWMutex
    cache map[string][]string
}

// Initialize a new local cache
func NewLocalCache() *LocalCache {
    return &LocalCache{
        cache: make(map[string][]string),
    }
}

// Fetches data from the cache, returns nil if not found
func (lc *LocalCache) Get(key string) ([]string, bool) {
    lc.mu.RLock()
    value, ok := lc.cache[key]
    lc.mu.RUnlock()
    // if ok {
    //     log.Printf("Cache hit for key: %s", key)
    // } else {
    //     log.Printf("Cache miss for key: %s", key)
    // }
    return value, ok
}

// Sets data in the cache
func (lc *LocalCache) Set(key string, value []string) {
    lc.mu.Lock()
    lc.cache[key] = value
    lc.mu.Unlock()
    // log.Printf("Set cache for key: %s", key)
}

// Depth-Limit Search
func DLS(depth int, startLink, goalLink string, currentDepth int, path []string, visited map[string]bool, cache *LocalCache) ([]string, error) {
    if currentDepth > depth {
        return nil, fmt.Errorf("reached maximum depth at depth %d", depth)
    }

    visited[startLink] = true
    // log.Printf("Visiting: %s", startLink)

    links, ok := cache.Get(startLink)
    if !ok {
        var err error
        links, err = fetchLinks(startLink)
        if err != nil {
            return nil, fmt.Errorf("error fetching links at title %s: %v", startLink, err)
        }
        cache.Set(startLink, links)
    }

    for _, link := range links {
        if !visited[link] {
            // log.Printf("Visiting new link: %s, Total links visited: %d, At depth : %d", link, len(visited), depth)
            visited[link] = true
        } else {
            // log.Printf("visit alreaddy link : %s", link)
            continue
        }

        if strings.EqualFold(link, goalLink) {
            return append(path, link), nil
        }

        new_path := append([]string{}, path...)
        new_path = append(new_path, link)
        result, err := DLS(depth, link, goalLink, currentDepth+1, new_path, visited, cache)
        if err == nil {
            return result, nil
        }
    }

    return nil, fmt.Errorf("goal not found at depth %d", depth)
}

func IDS(startLink, goalLink string) ([]string, error, int) {
    var i int = 0
    cache := NewLocalCache()
    
    for {
        visited := make(map[string]bool)
        result, err := DLS(i, startLink, goalLink, 0, []string{startLink}, visited, cache)
        if err == nil {
            return result, nil, len(visited)
        }
        log.Printf("No result at depth %d, error: %v", i, err)
        i++
        if i > 20 {
            break
        }
    }
    return nil, fmt.Errorf("goal not found after depth %d", i), 0       
}

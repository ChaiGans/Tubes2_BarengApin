package main

import (
	"fmt"
	"log"
	"strings"
)

// Depth-Limit Search
func DLS(depth int, starting_title_link string, goal_title_link string, current_depth int, temp_string_save []string, visited map[string]bool, total_link_iterate *int) ([]string, error) {
    if current_depth > depth {
        return nil, fmt.Errorf("reached maximum depth at depth %d", depth)
    }

    if strings.EqualFold(starting_title_link, goal_title_link) {
        return temp_string_save, nil
    }

    visited[starting_title_link] = true

    links, err := fetchLinks(starting_title_link)
    if err != nil {
        return nil, fmt.Errorf("error fetching links at title %s: %v", starting_title_link, err)
    }

    for _, link := range links {
        if !visited[link] {
            *total_link_iterate++
            visited[link] = true
        }

        new_path := append(temp_string_save, link)
        result, err := DLS(depth, link, goal_title_link, current_depth+1, new_path, visited, total_link_iterate)
        if err == nil {
            return result, nil
        }
    }

    return nil, fmt.Errorf("goal not found at depth %d", depth)
}

func IDS(starting_title_link string, goal_title_link string) ([]string, error, int) {
    var i int = 0
    var iteration_number int = 0
    visited := make(map[string]bool)

    for {
        result, err := DLS(i, starting_title_link, goal_title_link, 0, []string{starting_title_link}, visited, &iteration_number)
        if err == nil {
            return result, nil, iteration_number
        }
        log.Printf("No result at depth %d, error: %v", i, err)
        i++
        if i > 20 {
            break
        }
    }
    return nil, fmt.Errorf("goal not found after depth %d", i), iteration_number
}
package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

func extractTitleFromURL(wikiURL string) (string, error) {
    parsedURL, err := url.Parse(wikiURL)
    if err != nil {
        return "", err
    }

    pathSegments := strings.Split(parsedURL.Path, "/")
    if len(pathSegments) < 2 {
        return "", fmt.Errorf("invalid Wikipedia URL format")
    }
    articleTitleEncoded := pathSegments[len(pathSegments)-1]

    articleTitle, err := url.QueryUnescape(articleTitleEncoded)
    if err != nil {
        return "", err
    }
    return strings.ReplaceAll(articleTitle, "_", " "), nil
}

func DFS(depth int, starting_title string, goal_title string, current_depth int, temp_string_save []string) ([]string, error) {
    if current_depth > depth {
        return nil, fmt.Errorf("reached maximum depth")
    }

    links, err := fetchLinks(starting_title)
    if err != nil {
        return nil, fmt.Errorf("error fetching links: %v", err)
    }

    for _, link := range links {
        title, err := extractTitleFromURL(link)

        if strings.EqualFold(title, goal_title) {
            return append(temp_string_save, link), nil
        }

        // Append the current title to the path and recurse
        new_path := append(temp_string_save, link)
        result, err := DFS(depth, title, goal_title, current_depth+1, new_path)
        if err == nil {
            return result, nil
        }
    }

    return nil, fmt.Errorf("goal not found")
}

func main() {
	start := time.Now()

	fmt.Println("Backend is running")
	titleToSearch := "Joko Widodo"
	goalSearch := "Gibran Rakabuming"

	result, err := DFS(3, titleToSearch, goalSearch, 0, []string{})
	if err != nil {
		fmt.Print(err)
	} else {
		for _, link := range result {
			fmt.Println(link)
		}
	}

	elapsed := time.Since(start)
	log.Printf("Execution time took %d ms", elapsed.Milliseconds())

	// Call the fetchLinks function and handle the result
	// links, err := fetchLinks(titleToSearch)
	// if err != nil {
	// 	log.Fatalf("Error fetching links: %v", err)
	// }

	// fmt.Println("Links found:")
	// for _, link := range links {
	// 	fmt.Println(link)
	// }
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello from Backend!")
    // })
    // http.ListenAndServe(":8080", nil)
}

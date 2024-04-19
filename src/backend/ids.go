package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// func extractTitleFromURL(wikiURL string) (string, error) {
//     parsedURL, err := url.Parse(wikiURL)
//     if err != nil {
//         return "", err
//     }

//     pathSegments := strings.Split(parsedURL.Path, "/")
//     if len(pathSegments) < 2 {
//         return "", fmt.Errorf("invalid Wikipedia URL format")
//     }
//     articleTitleEncoded := pathSegments[len(pathSegments)-1]

//     return strings.ReplaceAll(articleTitleEncoded, "_", " "), nil
// }

// Depth-Limit Search
func DLS(depth int, starting_title_link string, goal_title_link string, current_depth int, temp_string_save []string) ([]string, error) {
    if current_depth > depth {
        return nil, fmt.Errorf("reached maximum depth at depth %d", depth)
    }

    if strings.EqualFold(starting_title_link, goal_title_link) {
        return temp_string_save, nil
    }

    links, err := scrapeWikipediaLinksAsync(starting_title_link)
    if err != nil {
        return nil, fmt.Errorf("error fetching links at title %s: %v", starting_title_link, err)
    }

    for _, link := range links {
        // title, err := extractTitleFromURL(link)
        // if err != nil {
        //     log.Printf("Error extracting title from URL: %v", err)
        //     continue
        // }

        if strings.EqualFold(link, goal_title_link) {
            return append(temp_string_save, link), nil
        }

        new_path := append(temp_string_save, link)
        result, err := DLS(depth, link, goal_title_link, current_depth+1, new_path)
        if err == nil {
            return result, nil
        }
    }

    return nil, fmt.Errorf("goal not found at depth %d", depth)
}

func IDS(starting_title_link string, goal_title_link string) ([]string, error) {
    var i int = 0
    for {
        result, err := DLS(i, starting_title_link, goal_title_link, 0, []string{starting_title_link})
        if err == nil {
            return result, nil
        }
        log.Printf("No result at depth %d, error: %v", i, err)
        i++
        if i > 20 {
            break
        }
    }
    return nil, fmt.Errorf("goal not found after depth %d", i)
}

func main() {
	start := time.Now()

	fmt.Println("Backend is running")
	titleToSearch := "Joko Widodo"
	goalSearch := "Pornography"

    title_to_search_link := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(titleToSearch, " ", "_")
    goal_to_search_link := "https://en.wikipedia.org/wiki/" + strings.ReplaceAll(goalSearch, " ", "_")

	result, err := IDS(title_to_search_link, goal_to_search_link);
	if err != nil {
		fmt.Print(err)
	} else {
		for _, link := range result {
			fmt.Println(link)
		}
	}

	elapsed := time.Since(start)
	log.Printf("Execution time took %d ms", elapsed.Milliseconds())
}


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
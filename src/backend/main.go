package main

import (
	"fmt"
	"log"
)
func main() {
	fmt.Println("backend is running")
    titleToSearch := "Joko Widodo" // Correctly initialize the variable

	// Call the fetchLinks function and handle the result
	links, err := fetchLinks(titleToSearch)
	if err != nil {
		log.Fatalf("Error fetching links: %v", err)
	}

	fmt.Println("Links found:")
	for _, link := range links {
		fmt.Println(link)
	}
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello from Backend!")
    // })
    // http.ListenAndServe(":8080", nil)
}

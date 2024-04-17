package main

import (
	"fmt"
	"net/http"
)
func main() {
	fmt.Println("backend is running")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Backend!")
    })
    http.ListenAndServe(":8080", nil)
}

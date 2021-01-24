package main

import (
	"fmt"
	"log"
	"net/http"
)

// placeholder front page
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, it's me.")
}

func main() {
	http.HandleFunc("/", handler) // placeholder front page
	log.Fatal(http.ListenAndServe(":8080", nil))
}

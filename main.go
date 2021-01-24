package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler) // placeholder front page
	http.ListenAndServe(port, nil)
}

// placeholder front page
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, it's me."))
}

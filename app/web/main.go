package main

import (
	// "html/template"
	"log"
	"net/http"
	"os"
	// "path/filepath"
)

func main() {
	// mux treats "/" like catch-all "/foo" also shows home()
	// http.HandleFunc() uses DefaultServeMux, which is a global variable
	// use locally scoped mux for security
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/code", showCode)
	mux.HandleFunc("/eat", showEat)
	mux.HandleFunc("/sleep", showSleep)

	// set dynamic port number for heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set.")
	}

	log.Println("Starting server on :"+port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}

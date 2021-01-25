package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday/v2"
)

// placeholder front page
func handler(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("markdown/index.md")
	if err != nil {
		log.Fatal(err)
	}
	md := blackfriday.Run(input)
	w.Write(md)
}

// nest blog.html into layout.html
func nestTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	// sanitize untrusted user request
	// [TODO] further sanitise before joining untrusted input
	// to prevent directory traversal attack
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func main() {
	fs := http.FileServer(http.Dir("./blogs"))
	http.Handle("/blogs/", http.StripPrefix("/blogs/", fs))

	http.HandleFunc("/", nestTemplate)

	// set dynamic port number for heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set.")
	}
	// deploy to port
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

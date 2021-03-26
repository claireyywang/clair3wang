// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// /*

// */
// func serveTemplate(w http.ResponseWriter, r *http.Request) {
// 	lp := filepath.Join("templates", "layout.html")
// 	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

// 	tmpl, _ := template.ParseFiles(lp, fp)
// 	tmpl.ExecuteTemplate(w, "layout", nil)
// }

// func main() {
// 	// set up home page
// 	http.HandleFunc("/", serveTemplate)

// 	// set dynamic port number for heroku deployment
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		log.Fatal("$PORT is not set.")
// 	}
// 	// deploy to port
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

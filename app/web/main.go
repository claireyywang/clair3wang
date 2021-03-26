package main

import (
	// "html/template"
	"fmt"
	"log"
	"net/http"
	"os"
	// "path/filepath"
)

// home Home page handler
// url "/"
func home(w http.ResponseWriter, r *http.Request) {
	// lp := filepath.Join("templates", "layout.html")
	// fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// tmpl, _ := template.ParseFiles(lp, fp)
	// tmpl.ExecuteTemplate(w, "layout", nil)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Claire"))
}

// showCode Code page handler
// url "/code"
func showCode(w http.ResponseWriter, r *http.Request) {
	// display a specific code project based on url query
	name := r.URL.Query().Get("project_name")
	if name != "" {
		fmt.Fprintf(w, "Display the chosen code project with name %s", name)
	} else {
		w.Write([]byte("Displaying all code projects..."))
	}
}

// showEat Eat page handler
// url "/eat"
func showEat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display Eat section ..."))
}

// showSleep Sleep page handler
// url "/sleep"
func showSleep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display Sleep section ..."))
}

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

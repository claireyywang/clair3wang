package main 

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

// home Home page handler
// url "/"
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// create a slice of template files
	// file path either needs to be relative to current work dir
	// or an abosolute path 
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// use template.ParseFiles() to read tmpl files
	// the slice of template files are passed as a variadic parameter
	// meaning there is no set number of files in `files`
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// then execute the template set
	// last param represent any dynamic data we want to 
	// pass in, which is nil for now
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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

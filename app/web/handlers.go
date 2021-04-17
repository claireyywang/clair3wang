package main 

import (
	"net/http"
	"html/template"
)

// createPage create page with templates
func (app *application) createPage(w http.ResponseWriter, fileName string) {
	// create a slice of template files
	// file path either needs to be relative to current work dir
	// or an abosolute path 
	files := []string{
		fileName,
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// use template.ParseFiles() to read tmpl files
	// the slice of template files are passed as a variadic parameter
	// meaning there is no set number of files in `files`
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// home Home page handler
// url "/"
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// create and execute the templated page
	app.createPage(w, "./ui/html/home.page.tmpl")
}

// showCode Code page handler
// url "/code"
func (app *application) showWork(w http.ResponseWriter, r *http.Request) {
	app.createPage(w, "./ui/html/work.page.tmpl")
}

// showEat Eat page handler
// url "/eat"
func (app *application) showLife(w http.ResponseWriter, r *http.Request) {
	app.createPage(w, "./ui/html/life.page.tmpl")
}

// showSleep Sleep page handler
// url "/sleep"
func (app *application) showContact(w http.ResponseWriter, r *http.Request) {
	app.createPage(w, "./ui/html/contact.page.tmpl")
}

package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
	"github.com/russross/blackfriday"
)

// createPage create page with templates
func (app *application) createPage(w http.ResponseWriter, fileName string) {
	// create a slice of template files
	// file path either needs to be relative to current work dir
	// or an abosolute path 
	files := []string{
		fileName,
		"./ui/html/base.layout.tmpl.html",
		"./ui/html/footer.partial.tmpl.html",
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

type Page struct {
	Body []byte
}

func markdownHelper(body []byte) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon(body))
}

func (app *application) renderMarkdown(fileName, tmplName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			"./ui/html/markdown.page.tmpl.html",
			tmplName,
			"./ui/html/base.layout.tmpl.html",
			"./ui/html/footer.partial.tmpl.html",
		}
		name := path.Base((files[0]))
	
		md, err := ioutil.ReadFile(fileName)
		p := &Page{Body: md}
	
		ts := template.Must(template.New(name).Funcs(template.FuncMap{"markDown": markdownHelper}).ParseFiles(files...))
		err = ts.ExecuteTemplate(w, name, p)
		if err != nil {
			app.serverError(w, err)
		}
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
	app.createPage(w, "./ui/html/home.page.tmpl.html")
}

// showWork Work page handler
// url "/work"
func (app *application) showWork(w http.ResponseWriter, r *http.Request) {
	app.createPage(w, "./ui/html/work.page.tmpl.html")
}

// showLife Life page handler
// url "/life"
func (app *application) showLife(w http.ResponseWriter, r *http.Request) {
	app.createPage(w, "./ui/html/life.page.tmpl.html")
}

// showContact Contact page handler
// url "/contact"
func (app *application) showContact(w http.ResponseWriter, r *http.Request) {
	app.createPage(w, "./ui/html/contact.page.tmpl.html")
}

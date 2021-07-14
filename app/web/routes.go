package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// mux treats "/" like catch-all "/foo" also shows home()
	// http.HandleFunc() uses DefaultServeMux, which is a global variable
	// use locally scoped mux for security
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/work", app.showWork)
	mux.HandleFunc("/work/openrobotics", app.renderMarkdown(
		"./content/markdown/openrobotics.md",
		"./ui/html/work.page.html"))
	mux.HandleFunc("/work/qualcomm", app.renderMarkdown(
		"./content/markdown/qualcomm.md",
		"./ui/html/work.page.html"))
	mux.HandleFunc("/work/upenn", app.renderMarkdown(
		"./content/markdown/upenn.md",
		"./ui/html/work.page.html"))
	mux.HandleFunc("/work/brynmawr", app.renderMarkdown(
		"./content/markdown/brynmawr.md",
		"./ui/html/work.page.html"))
	mux.HandleFunc("/projects", app.showProjects)
	mux.HandleFunc("/projects/shapes", app.renderArt("./ui/html/shapes.page.html"))
	mux.HandleFunc("/projects/cubes", app.renderArt("./ui/html/cubes.page.html"))
	mux.HandleFunc("/contact", app.showContact)

	// create a file serve which serves files out of ./ui/static dir
	// path given to http.Dir is relative to project root
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
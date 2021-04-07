package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// mux treats "/" like catch-all "/foo" also shows home()
	// http.HandleFunc() uses DefaultServeMux, which is a global variable
	// use locally scoped mux for security
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/code", app.showCode)
	mux.HandleFunc("/eat", app.showEat)
	mux.HandleFunc("/sleep", app.showSleep)

	// create a file serve which serves files out of ./ui/static dir
	// path given to http.Dir is relative to project root
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
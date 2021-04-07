package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError helper writes error msg and stack trace to errorLog
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
}

// clientError helper sends specific status
func (app *application) clientError(w http.ResponseWriter, status int) {
	// StatusText translates machine status to human readable text
	http.Error(w, http.StatusText(status), status)
}

// notFound helper that wraps around clientError which sends 404 Not Found
// using constant to eliminate hardcoded int 
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

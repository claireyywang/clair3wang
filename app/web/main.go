package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application Struct to hold the app dependencies
// include the custom loggers to be used in handlers
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	// leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize a new instance of application with the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	// mux treats "/" like catch-all "/foo" also shows home()
	// http.HandleFunc() uses DefaultServeMux, which is a global variable
	// use locally scoped mux for security
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/code", app.showCode)
	mux.HandleFunc("/eat", showEat)
	mux.HandleFunc("/sleep", showSleep)

	// create a file serve which serves files out of ./ui/static dir
	// path given to http.Dir is relative to project root
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// set dynamic port number for heroku deployment
	port := ":"+os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set.")
	}

	// initiate a new http.Server struct to use errorlog
	addr := flag.String("addr", port, "HTTP network address")
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog, 
		Handler: mux,
	}

	infoLog.Printf("Starting server on :"+port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

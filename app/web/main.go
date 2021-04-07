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

// main parse runtin config, establish dependencies, run HTTP server
func main() {
	// leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize a new instance of application with the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	

	// set dynamic port number for heroku deployment
	port := ":"+os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT is not set.")
	}

	// initiate a new http.Server struct to use errorlog
	addr := flag.String("addr", port, "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog, 
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

package main

import (
	"crypto/tls"
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

	// initiate a config struct to hold non-default settings
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		// restrict to strong modern cipher suites
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// initiate a new http.Server struct to use errorlog
	addr := flag.String("addr", port, "HTTP network address")
	flag.Parse()
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog, 
		Handler: app.routes(),
		TLSConfig: tlsConfig,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

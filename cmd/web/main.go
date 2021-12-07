package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.LUTC)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.”
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Shortcut, without config: err := http.ListenAndServe(*addr, mux)
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}

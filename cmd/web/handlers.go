package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// Home page
func home(w http.ResponseWriter, r *http.Request) {
	// Avoid the default catch-all behavior of "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

// Shows a snippet for a given Id
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		/*
			w.WriteHeader(405)
			w.Write([]byte("Method Not Allowed"))*/
		http.Error(w, "Method Not Allowed", 405) //The http.Error Shortcut

		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))
	//w.Write([]byte("Create a new snippet..."))
}

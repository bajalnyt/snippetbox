package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Avoid the default catch-all behavior of "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
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

// Shows a snippet for a given Id
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Create a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))
		//http.Error(w, "Method Not Allowed", 405) //The http.Error Shortcut
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"name":"Alex"}`))
	//w.Write([]byte("Create a new snippet..."))
}

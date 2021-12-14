package main

import (
	"bajal/snippetbox/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})

}

// Shows a snippet for a given Id
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	fmt.Println(id)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})

}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

// Create a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//w.Header().Set("content-type", "application/json")
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	validation_errors := make(map[string]string)
	if strings.TrimSpace(title) == "" {
		validation_errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		validation_errors["title"] = "This field is too long. Max is 1000"
	}

	if strings.TrimSpace(content) == "" {
		validation_errors["content"] = "Content cannot be empty"
	}

	if strings.TrimSpace(expires) == "" {
		validation_errors["expires"] = "Expiration cannot be empty"
	} else if expires != "365" && expires != "7" && expires != "1" {
		validation_errors["expires"] = "Expiration field invalid"
	}

	// In case of errors, redisplay page
	if len(validation_errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: validation_errors,
			FormData:   r.PostForm,
		})
	}
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

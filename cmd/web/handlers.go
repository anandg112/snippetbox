package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"strconv"

	"snippetbox.anandgautam.io/internal/models"
)

// change the signature of the home handler so it is defined as a method against *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)

	// if err != nil {
	// 	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	// 	app.serverError(w, r, err)
	// }

}

// Add a snippetView handler function.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

// Add a snippetCreate handler function.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	title := "snail"
	content := "snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, strconv.Itoa(expires))
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

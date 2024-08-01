package main

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"

	"github.com/molnarjani/GoForIt/models"
	"github.com/molnarjani/GoForIt/templates"
	"github.com/molnarjani/GoForIt/templates/pages"
)

func handleBadRequest(w http.ResponseWriter, r *http.Request) {
	slog.Error("request API", "method", r.Method, "status", http.StatusBadRequest, "path", r.URL.Path)
	w.WriteHeader(http.StatusBadRequest)
}

func renderTemplate(r *http.Request, w http.ResponseWriter, template templ.Component) {
	if err := htmx.NewResponse().RenderTempl(r.Context(), w, template); err != nil {
		slog.Error("render template", "method", r.Method, "status", http.StatusInternalServerError, "path", r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// indexViewHandler handles a view for the index page.
func indexViewHandler(w http.ResponseWriter, r *http.Request, ts *models.TodoStore) {
	// Define template meta tags.
	metaTags := pages.MetaTags(
		"goforit, todo, app, htmx, server, side, rendering",                                              // define meta keywords
		"This is an example todo app using server side rendering with HTMX, TailwindCSS, and Go backend", // define meta description
	)

	// Define template body content.
	bodyContent := pages.BodyContent(
		ts.List(),
	)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Go For It!",
		metaTags,
		bodyContent,
	)

	renderTemplate(r, w, indexTemplate)
}

func addTodoAPIHandler(w http.ResponseWriter, r *http.Request, ts *models.TodoStore) {
	if !htmx.IsHTMX(r) {
		handleBadRequest(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	if err := r.ParseForm(); err != nil {
		handleBadRequest(w, r)
		return
	}

	title := r.FormValue("title")

	ts.Add(models.Todo{
		Id:    strconv.Itoa(len(ts.List())),
		Title: title,
		Done:  false,
	})

	todosListTemplate := pages.TodosContent(
		ts.List(),
	)

	renderTemplate(r, w, todosListTemplate)
}

func updateTodoAPIHandler(w http.ResponseWriter, r *http.Request, ts *models.TodoStore) {
	if !htmx.IsHTMX(r) {
		handleBadRequest(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	if err := r.ParseForm(); err != nil {
		handleBadRequest(w, r)
		return
	}

	todoId := r.PathValue("id")
	_, todo, _ := ts.Get(todoId)
	done := r.FormValue("done") == "on"
	ts.Update(todoId, models.Todo{
		Id:    todo.Id,
		Done:  done,
		Title: todo.Title,
	})
}

func deleteTodoAPIHandler(w http.ResponseWriter, r *http.Request, ts *models.TodoStore) {
	if !htmx.IsHTMX(r) {
		handleBadRequest(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	if err := r.ParseForm(); err != nil {
		handleBadRequest(w, r)
		return
	}

	todoId := r.PathValue("id")
	slog.Info("todoId", "todoId", todoId)
	ts.Delete(todoId)

	todosListTemplate := pages.TodosContent(
		ts.List(),
	)

	renderTemplate(r, w, todosListTemplate)
}

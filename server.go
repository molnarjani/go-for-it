package main

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	gowebly "github.com/gowebly/helpers"
	"github.com/molnarjani/GoForIt/models"
)

//go:embed all:static
var static embed.FS

func injectTodoStore(ts *models.TodoStore, handler func(w http.ResponseWriter, r *http.Request, ts *models.TodoStore)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, ts)
	}
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	TodoStore := models.NewTodoStore()

	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "7000"))
	if err != nil {
		return err
	}

	// Handle static files from the embed FS (with a custom handler).
	http.Handle("GET /static/", gowebly.StaticFileServerHandler(http.FS(static)))

	// Handle index page view.
	http.HandleFunc("GET /", injectTodoStore(TodoStore, indexViewHandler))

	// Handle API endpoints with logging middleware.
	http.Handle("POST /api/add-todo", loggingMiddleware(injectTodoStore(TodoStore, addTodoAPIHandler)))
	http.Handle("PATCH /api/update-todo/{id}", loggingMiddleware(injectTodoStore(TodoStore, updateTodoAPIHandler)))
	http.Handle("DELETE /api/delete-todo/{id}", loggingMiddleware(injectTodoStore(TodoStore, deleteTodoAPIHandler)))

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}

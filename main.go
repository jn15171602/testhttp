package main

import (
	"example/hellohttp/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.View().Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", r)
}

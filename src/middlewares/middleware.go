package middlewares

import (
	"net/http"
	"olist-challenge/src/entities/book"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Middleware() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Group(book.Routes)

	return r
}

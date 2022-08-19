package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Middleware() http.Handler {
	r := chi.NewRouter()

	return r
}

package book

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Get("/", GetAllController)
	r.Get("/{id}", GetOneController)
	r.Post("/", CreateOneController)
	r.Put("/{id}", UpdateOneController)
	r.Delete("/{id}", DeleteOneController)

}

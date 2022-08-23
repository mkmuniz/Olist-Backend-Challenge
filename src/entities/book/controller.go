package book

import (
	"encoding/json"
	"net/http"

	httperror "olist-challenge/src/utils/errors/http"

	"github.com/go-chi/chi/v5"
)

func GetAllController(w http.ResponseWriter, r *http.Request) {
	res, err := GetAllService()

	httperror.RequestError(w, r, res, err)
}

func GetOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := GetOneService(id)

	httperror.RequestError(w, r, res, err)
}

func CreateOneController(w http.ResponseWriter, r *http.Request) {
	var body Book

	err := json.NewDecoder(r.Body).Decode(&body)

	httperror.DecodeError(w, r, err)

	res, err := CreateOneService(body)

	httperror.RequestError(w, r, res, err)
}

func UpdateOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body Book

	err := json.NewDecoder(r.Body).Decode(&body)

	httperror.DecodeError(w, r, err)

	res, err := UpdateOneService(id, body)

	httperror.RequestError(w, r, res, err)
}

func DeleteOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := DeleteOneService(id)

	httperror.RequestError(w, r, res, err)
}

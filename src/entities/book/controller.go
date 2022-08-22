package book

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var resp map[string]any

func GetAllController(w http.ResponseWriter, r *http.Request) {
	res, err := GetAllService()

	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on request",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Success on get user with id",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := GetOneService(id)

	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error on get user with id",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Success on get user with id",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func CreateOneController(w http.ResponseWriter, r *http.Request) {
	var body Book

	err := json.NewDecoder(r.Body).Decode(&body)

	res, err := CreateOneService(body)

	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on request",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Success on create book",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func UpdateOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var body Book

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on request",
			"data":    body,
		}
	}

	res, err := UpdateOneService(id, body)

	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error on update user",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Success on update user",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DeleteOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := DeleteOneService(id)

	if err != nil {
		resp = map[string]any{
			"status":  404,
			"message": "Error on delete user with id",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  200,
			"message": "Success on delete user with id",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

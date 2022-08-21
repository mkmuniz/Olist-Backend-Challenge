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
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetOneController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := GetOneService(id)

	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on request",
			"data":    res,
		}
	}

	if err != nil {
		resp = map[string]any{
			"status":  500,
			"message": "Error on get user with id",
			"data":    res,
		}
	} else {
		resp = map[string]any{
			"status":  404,
			"message": "Success on get user with id",
			"data":    res,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func CreateOneController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetOneController"))
}

func UpdateOneController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetOneController"))
}

func DeleteOneController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetOneController"))
}

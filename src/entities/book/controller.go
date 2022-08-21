package book

import (
	"encoding/json"
	"net/http"
)

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
	w.Write([]byte("GetOneController"))
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

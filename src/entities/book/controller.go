package book

import "net/http"

func GetAllController(w http.ResponseWriter, r *http.Request) {
	GetAllService()
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

package server

import (
	"net/http"
	"olist-challenge/configs"
	"olist-challenge/src/middlewares"
)

func Run() {

	configs.Load()
	http.ListenAndServe(":8080", middlewares.Middleware())
}

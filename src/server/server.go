package server

import (
	"log"
	"net/http"
	"olist-challenge/configs"
	"olist-challenge/src/db"
	"olist-challenge/src/middlewares"
)

func Run() {
	configs.Load()
	db.OpenConnection()
	log.Printf("API is running at port %v", configs.GetAPIConfig())
	http.ListenAndServe(configs.GetAPIConfig(), middlewares.Middleware())
}

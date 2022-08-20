package book

import (
	"log"
	"olist-challenge/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllService() *mongo.SingleResult {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Panic(err)
	}

	result := conn.Database.Collection("books").FindOne(conn.Context, bson.D{})

	defer conn.Client.Disconnect(conn.Context)

	return result
}

func GetOneService() {

}

func CreateOneService() {

}

func UpdateOneService() {

}

func DeleteOneService() {

}

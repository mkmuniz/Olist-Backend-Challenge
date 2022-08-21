package book

import (
	"log"
	"olist-challenge/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllService() (books bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Panic(err)
	}

	result, err := conn.Database.Collection("books").Find(conn.Context, bson.M{})

	if err != nil {
		log.Panic(err)
	}

	defer result.Close(conn.Context)

	if result.Next(conn.Context) {
		var books bson.M
		if err = result.Decode(&books); err != nil {
			log.Panic(err)
		}
		return books, err
	}
	return books, err
}

func GetOneService(id string) (books bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return books, err
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return books, err
	}

	res := conn.Database.Collection("books").FindOne(conn.Context, bson.M{"_id": objectId})

	res.Decode(&books)

	return books, err
}

func CreateOneService() {

}

func UpdateOneService() {

}

func DeleteOneService() {

}

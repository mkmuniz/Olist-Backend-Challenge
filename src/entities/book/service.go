package book

import (
	"log"
	"olist-challenge/src/db"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllService() (books bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Panic(err)
	}

	result, err := conn.Database.Collection("books").Find(conn.Context, bson.M{})

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

func GetOneService(body Book) (books bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Panic(err)
	}

	res := conn.Database.Collection("books").FindOne(conn.Context, bson.M{"_id": body.id})

	res.Decode(&books)

	return books, err
}

func CreateOneService() {

}

func UpdateOneService() {

}

func DeleteOneService() {

}

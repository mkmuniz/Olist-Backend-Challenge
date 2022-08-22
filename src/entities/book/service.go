package book

import (
	"log"
	"olist-challenge/src/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllService() (books []bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		log.Panic(err)
	}

	result, err := conn.Database.Collection("books").Find(conn.Context, bson.M{})

	if err != nil {
		log.Panic(err)
	}

	result.All(conn.Context, &books)

	defer result.Close(conn.Context)

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

func CreateOneService(body Book) (id interface{}, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return id, err
	}

	res, err := conn.Database.Collection("books").InsertOne(conn.Context, body)

	id = res.InsertedID

	if err != nil {
		return id, err
	}

	return id, err

}

func UpdateOneService(id string, body Book) (books bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return books, err
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return books, err
	}

	res := conn.Database.Collection("books").FindOneAndUpdate(conn.Context, bson.M{"_id": objectId}, bson.M{"$set": body})

	res.Decode(&books)

	return books, err
}

func DeleteOneService(id string) (books bson.M, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return books, err
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return books, err
	}

	res := conn.Database.Collection("books").FindOneAndDelete(conn.Context, bson.M{"_id": objectId})

	res.Decode(&books)

	return books, err
}

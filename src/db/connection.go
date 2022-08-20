package db

import (
	"context"
	"log"
	"olist-challenge/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (conn Connection, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(configs.GetDBConfig().URL))

	if err != nil {
		log.Panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Panic(err)
	}

	err = client.Ping(ctx, nil)

	conn.Client = client
	conn.Database = client.Database(configs.GetDBConfig().Name)
	conn.Context = ctx

	return conn, err
}

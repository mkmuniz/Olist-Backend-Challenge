package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Connection struct {
	Client   *mongo.Client
	Database *mongo.Database
	Context  context.Context
}

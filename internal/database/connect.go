package database

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(DB_URL string) (*mongo.Client, error) {
	if DB_URL == "" {
		return nil, errors.New("empty DB_URL")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(DB_URL).SetServerAPIOptions(serverAPI).SetMaxPoolSize(50)
	client, err := mongo.Connect(opts)
	return client, err
}

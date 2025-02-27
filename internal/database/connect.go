package database

import (
	"context"
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

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client, err
}

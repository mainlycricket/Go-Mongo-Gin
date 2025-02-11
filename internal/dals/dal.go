package dals

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type dBContext struct {
	database   *mongo.Database
	collection *mongo.Collection
}

type dal[T any] struct {
	*dBContext
}

func (d *dal[T]) InsertOne(ctx context.Context, item *T) (bson.ObjectID, error) {
	result, err := d.collection.InsertOne(ctx, item)
	if err != nil {
		return bson.NilObjectID, err
	}

	return result.InsertedID.(bson.ObjectID), nil
}

func (d *dal[T]) ReadAll(ctx context.Context) ([]T, error) {
	cursor, err := d.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var result []T
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, err
}

func (d *dal[T]) ReadById(ctx context.Context, id bson.ObjectID) (*T, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var item T
	if err := result.Decode(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (d *dal[T]) DeleteById(ctx context.Context, id bson.ObjectID) error {
	filters := bson.D{{Key: "_id", Value: id}}
	result, err := d.collection.DeleteOne(ctx, filters)
	if err != nil {
		return err
	}

	if result.DeletedCount != -1 {
		return errors.New("failed to delete document")
	}

	return nil
}

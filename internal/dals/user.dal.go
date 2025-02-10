package dals

import (
	"context"

	"github.com/mainlycricket/go-mongo/internal/database/models"
	"github.com/mainlycricket/go-mongo/internal/dtos/responses"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IUserDal interface {
	IBaseDal[models.User]
	ReadAll(ctx context.Context) ([]responses.AllUserResponse, error)
}

type userDal struct {
	dal[models.User]
}

func NewUserDal(database *mongo.Database) IUserDal {
	userDal := userDal{
		dal: dal[models.User]{
			dBContext: &dBContext{
				database:   database,
				collection: database.Collection("users"),
			},
		},
	}

	return &userDal
}

func (ud *userDal) ReadAll(ctx context.Context) ([]responses.AllUserResponse, error) {
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})

	cursor, err := ud.collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var users []responses.AllUserResponse
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

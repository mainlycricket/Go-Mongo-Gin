package services

import (
	"context"

	"github.com/mainlycricket/go-mongo/internal/dals"
	"github.com/mainlycricket/go-mongo/internal/database/models"
	"github.com/mainlycricket/go-mongo/internal/dtos/responses"
	"github.com/mainlycricket/go-mongo/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserDal interface {
	InsertOne(ctx context.Context, item *models.User) (bson.ObjectID, error)
	ReadById(ctx context.Context, id bson.ObjectID) (*models.User, error)
	DeleteById(ctx context.Context, id bson.ObjectID) error
	ReadAll(ctx context.Context) ([]responses.AllUserResponse, error)
}

type UserService struct {
	dal IUserDal
}

func NewUserService(database *mongo.Database) *UserService {
	userService := UserService{
		dal: dals.NewUserDal(database),
	}

	return &userService
}

func (us *UserService) InsertUser(ctx context.Context, user *models.User) (bson.ObjectID, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return bson.NilObjectID, err
	}

	user.Password = hashedPassword
	return us.dal.InsertOne(ctx, user)
}

func (us *UserService) ReadAll(ctx context.Context) ([]responses.AllUserResponse, error) {
	return us.dal.ReadAll(ctx)
}

func (us *UserService) ReadById(ctx context.Context, id bson.ObjectID) (*models.User, error) {
	return us.dal.ReadById(ctx, id)
}

func (us *UserService) DeleteById(ctx context.Context, id bson.ObjectID) error {
	return us.dal.DeleteById(ctx, id)
}

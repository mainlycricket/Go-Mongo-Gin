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

type IUserService interface {
	InsertUser(ctx context.Context, user *models.User) (bson.ObjectID, error)
	ReadById(ctx context.Context, id bson.ObjectID) (*models.User, error)
	ReadAll(ctx context.Context) ([]responses.AllUserResponse, error)
	DeleteById(ctx context.Context, id bson.ObjectID) error
}

type userService struct {
	dal dals.IUserDal
}

func NewUserService(database *mongo.Database) IUserService {
	userService := userService{
		dal: dals.NewUserDal(database),
	}

	return &userService
}

func (us *userService) InsertUser(ctx context.Context, user *models.User) (bson.ObjectID, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return bson.NilObjectID, err
	}

	user.Password = hashedPassword
	return us.dal.InsertOne(ctx, user)
}

func (us *userService) ReadAll(ctx context.Context) ([]responses.AllUserResponse, error) {
	return us.dal.ReadAll(ctx)
}

func (us *userService) ReadById(ctx context.Context, id bson.ObjectID) (*models.User, error) {
	return us.dal.ReadById(ctx, id)
}

func (us *userService) DeleteById(ctx context.Context, id bson.ObjectID) error {
	return us.dal.DeleteById(ctx, id)
}

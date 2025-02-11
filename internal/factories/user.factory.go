package factories

import (
	"context"
	"fmt"

	"github.com/mainlycricket/go-mongo/internal/database/models"
	"github.com/mainlycricket/go-mongo/internal/dtos/responses"
	"github.com/mainlycricket/go-mongo/internal/services"
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

type UserFactory struct {
	service IUserService
}

func NewUserFactory(database *mongo.Database) *UserFactory {
	userFactory := UserFactory{
		service: services.NewUserService(database),
	}

	return &userFactory
}

func (uf *UserFactory) InsertUser(ctx context.Context, user *models.User) (bson.ObjectID, error) {
	userId, err := uf.service.InsertUser(ctx, user)
	if err != nil {
		return bson.NilObjectID, err
	}

	otpEmailInput := utils.GetOtpEmailInput(user.Email)
	if err := utils.SendEmail(otpEmailInput); err != nil {
		if err := uf.service.DeleteById(ctx, userId); err != nil {

		}
		return bson.NilObjectID, fmt.Errorf(`failed to send OTP verification to %s`, user.Email)
	}

	return userId, nil
}

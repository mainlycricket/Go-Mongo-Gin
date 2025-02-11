package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mainlycricket/go-mongo/internal/database"
	"github.com/mainlycricket/go-mongo/internal/database/models"
	"github.com/mainlycricket/go-mongo/internal/dtos/responses"
	"github.com/mainlycricket/go-mongo/internal/factories"
	"github.com/mainlycricket/go-mongo/internal/services"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserService interface {
	InsertUser(ctx context.Context, user *models.User) (bson.ObjectID, error)
	ReadById(ctx context.Context, id bson.ObjectID) (*models.User, error)
	ReadAll(ctx context.Context) ([]responses.AllUserResponse, error)
	DeleteById(ctx context.Context, id bson.ObjectID) error
}

type IUserFactory interface {
	InsertUser(ctx context.Context, user *models.User) (bson.ObjectID, error)
}

type UserHandler struct {
	userService IUserService
	userFactory IUserFactory
}

func NewUserHandler(database *mongo.Database) *UserHandler {
	handler := UserHandler{
		userService: services.NewUserService(database),
		userFactory: factories.NewUserFactory(database),
	}

	return &handler
}

func (uh *UserHandler) CreateUser(g *gin.Context) {
	var user models.User
	var response responses.DefaultApiResponse

	if err := g.ShouldBindJSON(&user); err != nil {
		response.Message = "invalid request body"
		g.JSON(http.StatusBadRequest, response)
		return
	}

	if errors := user.Validate(); len(errors) > 0 {
		response.Message = "validated failed"
		response.Data = errors
		g.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := uh.userFactory.InsertUser(g.Request.Context(), &user)
	if err != nil {
		response.Message = err.Error()
		response.Data = nil
		g.JSON(database.GetHttpStatusByDbError(err), response)
		return
	}

	response.Successs = true
	response.Message = "user inserted successfully, check email for verification code"
	response.Data = responses.UserInsertionResponse{InsertedId: id.Hex()}
	g.JSON(http.StatusCreated, response)
}

func (uh *UserHandler) GetAllUsers(g *gin.Context) {
	var response responses.DefaultApiResponse

	users, err := uh.userService.ReadAll(g.Request.Context())
	if err != nil {
		response.Message = "failed to get all users"
		response.Data = err
		g.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Successs = true
	response.Message = "read all users successfully!"
	response.Data = users
	g.JSON(http.StatusOK, response)
}

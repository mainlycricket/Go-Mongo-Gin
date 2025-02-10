package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mainlycricket/go-mongo/internal/database"
	"github.com/mainlycricket/go-mongo/internal/database/models"
	"github.com/mainlycricket/go-mongo/internal/dtos/responses"
	"github.com/mainlycricket/go-mongo/internal/factories"
	"github.com/mainlycricket/go-mongo/internal/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserHandler interface {
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
}

func NewUserHandler(database *mongo.Database) IUserHandler {
	handler := userHandler{
		userService: services.NewUserService(database),
		userFactory: factories.NewUserFactory(database),
	}

	return &handler
}

type userHandler struct {
	userService services.IUserService
	userFactory factories.IUserFactory
}

func (uh *userHandler) CreateUser(g *gin.Context) {
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

func (uh *userHandler) GetAllUsers(g *gin.Context) {
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

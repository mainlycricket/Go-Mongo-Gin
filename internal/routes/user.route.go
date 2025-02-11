package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mainlycricket/go-mongo/internal/handlers"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IUserHandler interface {
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
}

func RegisterUserRoutes(r *gin.RouterGroup, database *mongo.Database) {
	var userHandler IUserHandler = handlers.NewUserHandler(database)

	userRoutes := r.Group("/users")
	userRoutes.GET("/", userHandler.GetAllUsers)
	userRoutes.POST("/", userHandler.CreateUser)
}

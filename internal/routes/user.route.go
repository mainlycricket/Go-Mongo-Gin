package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mainlycricket/go-mongo/internal/handlers"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterUserRoutes(r *gin.RouterGroup, database *mongo.Database) {
	userHandler := handlers.NewUserHandler(database)

	userRoutes := r.Group("/users")
	userRoutes.GET("/", userHandler.GetAllUsers)
	userRoutes.POST("/", userHandler.CreateUser)
}

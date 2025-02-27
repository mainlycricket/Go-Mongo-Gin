package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mainlycricket/go-mongo/internal/database"
	"github.com/mainlycricket/go-mongo/internal/dtos/responses"
	"github.com/mainlycricket/go-mongo/internal/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	database, err := initDB()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	ginEngine := gin.Default()
	ginEngine.SetTrustedProxies(nil)

	ginEngine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, responses.DefaultApiResponse{Successs: true, Message: "running!!"})
		return
	})

	apiRoutes := ginEngine.Group("/api")

	v1Routes := apiRoutes.Group("/v1")
	routes.RegisterUserRoutes(v1Routes, database)

	ginEngine.Run(":8080")
}

func initDB() (*mongo.Database, error) {
	client, err := database.Connect(os.Getenv("CLUSTER_URL"))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client.Database(os.Getenv("DB_NAME")), nil
}

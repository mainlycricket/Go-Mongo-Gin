package main

import (
	"context"
	"errors"
	"fmt"
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
	host, port := os.Getenv("DB_HOST"), os.Getenv("DB_PORT")

	if host == "" || port == "" {
		return nil, errors.New("missing db env vars")
	}

	username, password := os.Getenv("MONGO_INITDB_ROOT_USERNAME"), os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	var credentials string

	if username != "" && password != "" {
		credentials = fmt.Sprintf("%s:%s@", username, password)
	}

	DB_URL := fmt.Sprintf("mongodb://%s%s:%s", credentials, host, port)

	client, err := database.Connect(DB_URL)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client.Database(os.Getenv("DB_NAME")), nil
}

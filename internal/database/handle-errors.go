package database

import (
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetHttpStatusByDbError(err error) int {
	if mongo.IsDuplicateKeyError(err) {
		return http.StatusBadRequest
	}

	switch err {
	case mongo.ErrNoDocuments:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

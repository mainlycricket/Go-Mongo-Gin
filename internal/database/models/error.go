package models

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

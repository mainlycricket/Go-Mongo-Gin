package responses

type UserInsertionResponse struct {
	InsertedId string `json:"inserted_id"`
}

type AllUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

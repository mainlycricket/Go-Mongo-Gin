package models

import "slices"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u *User) Validate() []ValidationError {
	var errors []ValidationError

	if u == nil {
		errors = append(errors, ValidationError{Error: "user object not found"})
		return errors
	}

	if u.Name == "" {
		errors = append(errors, ValidationError{Field: "name", Error: "name is required"})
	}

	if u.Email == "" {
		errors = append(errors, ValidationError{Field: "email", Error: "email is required"})
	}

	// TODO email regex pending

	if u.Password == "" {
		errors = append(errors, ValidationError{Field: "passowrd", Error: "password is required"})
	}

	// TODO password constraints

	if !slices.Contains([]string{"admin", "student"}, u.Role) {
		errors = append(errors, ValidationError{Field: "role", Error: "invalid role"})
	}

	return errors
}

package users

import (
	"encoding/json"
	"net/http"
)

const MaxListLimit = 200

type GetUserRequest struct {
	ID string `json:"id"`
}

type ListUserRequest struct {
	Limit int    `json:"limit"`
	After string `json:"after"`
	Name  string `json:"name"`
}

type CreateUserRequest struct {
	User *User `json:"user"`
}

type UpdateUserRequest struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	EmailId     string `json:"email"`
	PhoneNumber string `json:"phone"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type UserResponseWrapper struct {
	User  *User   `json:"user,omitempty"`
	Users []*User `json:"users,omitempty"`
	Code  int     `json:"-"`
}

func (e *UserResponseWrapper) Json() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

func (e *UserResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}

	return e.Code
}

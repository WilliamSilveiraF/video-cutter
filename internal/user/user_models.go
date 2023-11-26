package user

import (
    "workflow-editor/internal/person"
)


type User struct {
	Email    string
	Password string
}

type RegisterUserRequest struct {
    User   User             `json:"user"`
    Person person.Person    `json:"person"`
}

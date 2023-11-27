package user

import (
    "workflow-editor/internal/person"
)


type User struct {
	Email    string	`json:"email"`
	Password string	`json:"password"`
}

type RegisterUserRequest struct {
    User   User             `json:"user"`
    Person person.Person    `json:"person"`
}

type UpdatePasswordRequest struct {
	Email		string	`json:"email"`
	OldPassword string	`json:"old_password"`
	NewPassword	string	`json:"new_password"`
}

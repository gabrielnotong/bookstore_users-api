package users

import (
	"github.com/gabrielnotong/bookstore_users-api/errors"
	validator2 "gopkg.in/go-playground/validator.v9"
)

type User struct {
	Id        int64  `json:"id" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
}

func (u *User) Validate() *errors.RestErr {
	err := validator2.New().StructExcept(u, "Id", "CreatedAt")
	if err != nil {
		return errors.NewBadRequestError("Invalid request body")
	}

	return nil
}

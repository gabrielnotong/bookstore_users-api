package users

import (
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"strings"
)

type User struct {
	Id        int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email")
	}

	return nil
}

package users

import (
	"fmt"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/formatting"
)

var (
	DB = make(map[int64]*User)
)

func (u *User) Find() *errors.RestErr {
	user := DB[u.Id]
	if user == nil {
		return errors.NewNotFoundError("user not found")
	}

	u.Id = user.Id
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.CreatedAt = user.CreatedAt

	return nil
}

func (u *User) Save() *errors.RestErr {
	current := DB[u.Id]
	if current != nil {
		if u.Email == DB[u.Id].Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.Id))
	}
	u.CreatedAt = formatting.DateNowString()
	DB[u.Id] = u
	return nil
}

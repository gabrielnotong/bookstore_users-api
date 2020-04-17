package services

import (
	"github.com/gabrielnotong/bookstore_users-api/domain/users"
	"github.com/gabrielnotong/bookstore_users-api/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	if err := users.NewValidator().Validate(&u); err != nil {
		return nil, err
	}

	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func FindUser(id int64) (*users.User, *errors.RestErr) {
	u := &users.User{Id: id}
	if err := u.Find(); err != nil {
		return nil, err
	}

	return u, nil
}

package services

import (
	"github.com/gabrielnotong/bookstore_users-api/domain/users"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/formatting"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	if err := users.NewValidator().Validate(&u); err != nil {
		return nil, err
	}

	u.Status = users.StatusActive
	u.CreatedAt = formatting.DateNowString()
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

func UpdateUser(u *users.User) (*users.User, *errors.RestErr) {
	cu, err := FindUser(u.Id)
	if err != nil {
		return nil, err
	}

	if err := users.NewValidator().Validate(u); err != nil {
		return nil, err
	}

	cu.FirstName = u.FirstName
	cu.LastName = u.LastName
	cu.Email = u.Email

	if err := cu.Update(); err != nil {
		return nil, err
	}

	return cu, nil
}

func DeleteUser(id int64) *errors.RestErr {
	cu := &users.User{Id: id}
	return cu.Delete()
}

func Search(status string) ([]*users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
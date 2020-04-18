package services

import (
	"github.com/gabrielnotong/bookstore_users-api/domain/users"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/formatting"
)

var (
	UsersService usersServiceInterface = &usersService{}
)
type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	FindUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(*users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

type usersService struct {}

func (us *usersService) CreateUser(u users.User) (*users.User, *errors.RestErr) {
	if err := users.Validator.Validate(&u); err != nil {
		return nil, err
	}

	u.Status = users.StatusActive
	u.CreatedAt = formatting.DateNowString()
	u.Password = formatting.Sha256(u.Password)
	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (us *usersService) FindUser(id int64) (*users.User, *errors.RestErr) {
	u := &users.User{Id: id}
	if err := u.Find(); err != nil {
		return nil, err
	}

	return u, nil
}

func (us *usersService) UpdateUser(u *users.User) (*users.User, *errors.RestErr) {
	cu, err := us.FindUser(u.Id)
	if err != nil {
		return nil, err
	}

	if err := users.Validator.Validate(u); err != nil {
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

func (us *usersService) DeleteUser(id int64) *errors.RestErr {
	cu := &users.User{Id: id}
	return cu.Delete()
}

func (us *usersService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

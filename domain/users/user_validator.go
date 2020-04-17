package users

import (
	"github.com/gabrielnotong/bookstore_users-api/errors"
	validator2 "gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	v2 *validator2.Validate
}

func NewValidator() *Validator {
	return &Validator{validator2.New()}
}

func (v *Validator) Validate(u *User) *errors.RestErr {
	err := v.v2.StructExcept(u)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
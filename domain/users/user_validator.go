package users

import (
	"github.com/gabrielnotong/bookstore_users-api/errors"
	validator2 "gopkg.in/go-playground/validator.v9"
	"strings"
)

var (
	Validator validatorInterface = &validator{validator2.New()}
)

type validator struct {
	v2 *validator2.Validate
}

type validatorInterface interface {
	Validate(*User) *errors.RestErr
}

func (v *validator) Validate(u *User) *errors.RestErr {
	err := v.v2.StructExcept(u, "Id", "CreatedAt")
	if err != nil {
		return errors.NewBadRequestError(formatErrorMsg(err))
	}

	return nil
}

func formatErrorMsg(err error) string {
	errStr := err.Error()

	switch {
	case strings.Contains(errStr, "'FirstName' failed on the 'required' tag"):
		return "first_name is required."
	case strings.Contains(errStr, "'LastName' failed on the 'required' tag"):
		return "last_name is required."
	case strings.Contains(errStr, "'Email' failed on the 'required' tag"):
		return "email is required."
	case strings.Contains(errStr, "'Password' failed on the 'required' tag"):
		return "password is required."
	}

	return errStr
}

package users

import (
	"github.com/gabrielnotong/bookstore_users-api/datasource/mysql/users_db"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/formatting"
)

var (
	DB = users_db.DB
)

func (u *User) Find() *errors.RestErr {
	row := DB.QueryRow("SELECT * FROM users WHERE id = $1", u.Id)
	if row == nil {
		return errors.NewNotFoundError("user not found")
	}

	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (u *User) Save() *errors.RestErr {
	u.CreatedAt = formatting.DateNowString()

	sqlStatement := `
	INSERT INTO users (first_name, last_name, email, created_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id int64

	err := DB.QueryRow(sqlStatement, u.FirstName, u.LastName, u.Email, u.CreatedAt).Scan(&id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	u.Id = id

	return nil
}

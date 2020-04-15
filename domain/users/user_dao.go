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
	stmt, err := DB.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		return errors.ParsePostgresError(err)
	}

	res := stmt.QueryRow(u.Id)

	err = res.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt)
	if err != nil {
		return errors.ParsePostgresError(err)
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
		return errors.ParsePostgresError(err)
	}

	u.Id = id

	return nil
}

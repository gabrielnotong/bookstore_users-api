package users

import (
	"fmt"
	"github.com/gabrielnotong/bookstore_users-api/datasource/mysql/users_db"
	"github.com/gabrielnotong/bookstore_users-api/errors"
	"github.com/gabrielnotong/bookstore_users-api/formatting"
)

var (
	DB = users_db.DB
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, status, password, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	querySelectUser = "SELECT * FROM users WHERE id = $1"
	queryUpdateUser = "UPDATE users SET first_name=$2, last_name=$3, email=$4 WHERE id=$1"
	queryDeleteUser = "DELETE FROM users WHERE id=$1"
	queryFindByStatus = "SELECT id, first_name, last_name, email, created_at, status FROM users WHERE status=$1"
)

func (u *User) Find() *errors.RestErr {
	stmt, err := DB.Prepare(querySelectUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow(u.Id)

	err = res.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt)
	if err != nil {
		return errors.ParsePostgresError(err)
	}

	return nil
}

func (u *User) Save() *errors.RestErr {
	var id int64

	u.CreatedAt = formatting.DateNowString()
	err := DB.QueryRow(
		queryInsertUser,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Status,
		u.Password,
		u.CreatedAt,
	).Scan(&id)
	if err != nil {
		return errors.ParsePostgresError(err)
	}

	u.Id = id

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := DB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id, u.FirstName, u.LastName, u.Email)
	if err != nil {
		return errors.ParsePostgresError(err)
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := DB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		return errors.ParsePostgresError(err)
	}

	return nil
}

func (u *User) FindByStatus(status string) ([]*User, *errors.RestErr) {
	stmt, err := DB.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.ParsePostgresError(err)
	}

	uu := make([]*User, 0)
	for rows.Next() {
		us := &User{}
		err := rows.Scan(&us.Id, &us.FirstName, &us.LastName, &us.Email, &us.CreatedAt, &us.Status)
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		uu = append(uu, us)
	}

	if len(uu) == 0 {
		return nil, errors.NewNotFoundError(
			fmt.Sprintf("No record matching the given status: %s", status),
		)
	}

	return uu, nil
}

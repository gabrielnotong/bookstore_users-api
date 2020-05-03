package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	dbUsername = "dbUsername"
	dbPassword = "dbPassword"
	dbHost     = "dbHost"
	dbSchema   = "dbSchema"
	dbDriver   = "dbDriver"
)

var (
	DB *sql.DB

	username = os.Getenv(dbUsername)
	password = os.Getenv(dbPassword)
	host     = os.Getenv(dbHost)
	schema   = os.Getenv(dbSchema)
	driver   = os.Getenv(dbDriver)
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf(
		"%s://%s:%s@%s/%s?sslmode=disable",
		driver, username, password, host, schema,
	)

	DB, err = sql.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	log.Println("You are logged in.")
}

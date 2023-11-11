package utility

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "postgres"
)

func ConnectionDB() (*sql.DB, error) {
	connectionPostgres := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)

	db, err := sql.Open("postgres", connectionPostgres)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("db connection failed...")
		return nil, err
	} else {
		fmt.Println("db connection successed...")
	}

	return db, nil
}

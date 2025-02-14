package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/harshdangi/go-rest-api-with-fiber/config"
)

var DB *sql.DB

func Connect() error {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsin str to int")
	}

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
}

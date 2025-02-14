package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/harshdangi/go-rest-api-with-fiber/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		return fmt.Errorf("Invalid port value: %s", p)
	}

	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))

	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}
	CreateProductTable()
	fmt.Println("Connection successful.")
	return nil
}

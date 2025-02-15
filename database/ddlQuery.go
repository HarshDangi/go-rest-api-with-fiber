package database

import (
	"fmt"
	"log"
)

func CreateProductTable() {
	_, err := DB.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		price integer,
		name text,
		description text,
		category text NOT NULL 
	)`, TableName))

	if err != nil {
		log.Fatal("Error creating the table: ", err)
	}
}

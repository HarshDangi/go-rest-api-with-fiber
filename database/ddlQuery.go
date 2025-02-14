package database

func CreateProductTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY
		price integer
		name text
		description text,
		category text NOT NULL 
	)`)
}

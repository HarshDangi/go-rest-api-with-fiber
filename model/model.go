package model

type Product struct {
	Name        string
	Description string
	Category    string
	Price       int
}

type Products []Product

package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harshdangi/go-rest-api-with-fiber/database"
	"github.com/harshdangi/go-rest-api-with-fiber/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query(fmt.Sprintf("SELECT name, description, category, price FROM %s ORDER BY name", database.TableName))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{} //iodmatic way to initialise a zero size slice
		err := rows.Scan(&product.Name, &product.Description, &product.Description, &product.Category)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result = append(result, product)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All product returned successfully",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return nil
}

func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	row, err := database.DB.Query(fmt.Sprintf("SELECT * FROM %s WHERE id = $1", database.TableName), id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})

	}
	defer row.Close()

	product := model.Product{}
	if !row.Next() {
		log.Println("No rows were returned!")
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product not found.",
		})
	}

	err = row.Scan(&id, &product.Price, &product.Name, &product.Description, &product.Category)
	if err != nil {
		return err
	}

	// return the product
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product fetched",
		"product": product,
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return nil
}

func CreateProduct(c *fiber.Ctx) error {
	product := model.Product{}

	if err := c.BodyParser(&product); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	var id = -1
	log.Println(database.TableName)
	err := database.DB.QueryRow(fmt.Sprintf("INSERT INTO %s (name, description, category, price) VALUES ($1, $2, $3, $4) RETURNING id", database.TableName), product.Name, product.Description, product.Category, product.Price).Scan(&id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	if err != nil {
		log.Fatal("Error retrieving last insert ID:", err)
		id = -1
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product created successfully",
		"product": product,
		"id":      id,
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := database.DB.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", database.TableName), id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	rowsDeleted, _ := res.RowsAffected()

	if rowsDeleted == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Product not found.",
		})
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "product deleted successfully",
	}); err != nil || rowsDeleted == 0 {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return nil
}

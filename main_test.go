package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/harshdangi/go-rest-api-with-fiber/database"
	"github.com/stretchr/testify/assert"
)

var app *fiber.App

func TestMain(m *testing.M) {
	if err := database.Connect("testing_products"); err != nil {
		log.Fatal(err)
	}
	app = setupServer()

	exitCode := m.Run()

	database.DB.Exec(fmt.Sprintf("DROP TABLE %s", database.TableName))
	os.Exit(exitCode)
}
func TestGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/", nil)

	res, err := app.Test(req, -1)

	assert.Nil(t, err)

	assert.Equal(t, 200, res.StatusCode)
}

func TestCreateProduct(t *testing.T) {

	payload, _ := json.Marshal(map[string]interface{}{"name": "Test product", "description": "Test description", "category": "Test category", "price": 56})

	req, _ := http.NewRequest("POST", "/api/", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	assert.Nil(t, err)

	msg, _ := io.ReadAll(res.Body)
	assert.Equal(t, 200, res.StatusCode, string(msg))
}

func TestGetSingleProduct(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/1", nil)

	res, err := app.Test(req, -1)

	assert.Nil(t, err)

	assert.Equal(t, 200, res.StatusCode)

	msg, _ := io.ReadAll(res.Body)
	log.Printf(string(msg))
}

func TestDeleteProduct(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/1", nil)

	res, err := app.Test(req, -1)

	assert.Nil(t, err)

	assert.Equal(t, 200, res.StatusCode)

	msg, _ := io.ReadAll(res.Body)
	log.Printf(string(msg))
}

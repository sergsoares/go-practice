package main

import (
	"crud-app/middleware"
	"crud-app/models"
)

func main() {
	// testInsertUser()
	testConnection()
}

func testConnection() {
	middleware.CreateConnection()
}

func testInsertUser() {
	user := models.User{
		ID:       1,
		Name:     "Sergio",
		Location: "Campo Grande",
		Age:      27,
	}

	middleware.InsertUser(user)
}

package main

import (
	"API-CRUD/initializers"
	"API-CRUD/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDatabase()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}

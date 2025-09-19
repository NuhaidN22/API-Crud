package main

import (
	"API-CRUD/controllers"
	"API-CRUD/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Init function called")
	initializers.LoadEnvVariable()
	initializers.ConnectDatabase()
}

func main() {

	fmt.Println("Hello, World!123")

	router := gin.Default()

	// router.GET("/", controllers.PostsCreate)
	router.POST("/posts", controllers.PostsCreate)

	router.GET("/posts", controllers.PostIndex)
	router.GET("/posts/:id", controllers.PostShow)

	router.PUT("/posts/:id", controllers.PostUpdate)

	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run() // listen and serve on 0.0.0.0:8080
}

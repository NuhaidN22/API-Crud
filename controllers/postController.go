package controllers

import (
	"API-CRUD/initializers"
	"API-CRUD/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//get data from req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//if no error

	//return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	//get posts

	var posts []models.Post
	initializers.DB.Find(&posts)

	///respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id from Url
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	///respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id from Url
	id := c.Param("id")

	//get the data from req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	//find the post we want to update
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body})

	//respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}
func PostDelete(c *gin.Context) {
	// get id from Url
	id := c.Param("id")

	//get the data from req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	//find the post we want to update
	var post models.Post
	initializers.DB.First(&post, id)

	//delete it
	initializers.DB.Delete(&post)
	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

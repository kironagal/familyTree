package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//get to retrive all the blog post
//receive specific post by id
//creating, updating new blogpost
//Deleting the blog post

// blog structure with below fields
type blog struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var db *gorm.DB

// Mock data to start with
var blogs = []blog{
	{Id: "01", Title: "TitleOne", Content: "This is content of TitleOne"},
	{Id: "02", Title: "TitleTwo", Content: "This is content of TitleTwo"},
	{Id: "03", Title: "TitleThree", Content: "This is content of TitleThree"},
	{Id: "04", Title: "TitleFour", Content: "This is content ofTitleFour"},
}

// Initializing the Database connection with Gorm Object-Relational Mapping
func initDB() {
	db, err := gorm.Open(sqlite.Open("blogPost.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect")
	}
	db.AutoMigrate(&blog{})
}

func main() {

	initDB()

	//importing gin package/ Gin Framework with httpMethods
	router := gin.Default()
	router.GET("/posts", getBlogPosts)
	router.GET("/posts:id", getBlogPostsById)
	router.POST("/posts", createBlogPost)
	router.PUT("/posts:id", updateBlogPost)
	router.DELETE("/posts:id", deleteBlogPost)

	//starting the server on localhost
	router.Run("localhost:8080")
}

// This helps in getting/retreving posts
func getBlogPosts(c *gin.Context) {
	var fblog []blog
	db.Find(&fblog)
	c.JSON(http.StatusOK, blogs)
}

// based on the id passed in header url we are getting blogpost by id
func getBlogPostsById(c *gin.Context) {
	//var newBlog blog
	// id := c.Param("Id")
	// db.First(&newBlog, id)
	// c.JSON(http.StatusOK, newBlog)
	id := c.Param("Id")
	for _, a := range blogs {
		if a.Id == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "blog not found"})
}

// creating a new blogpost which will be appended to the exiting blog{}
func createBlogPost(c *gin.Context) {
	var newBlog blog
	//BindJson to bind the new blog to the exiting list of blog
	if err := c.BindJSON(&newBlog); err != nil {
		return
	}
	blogs = append(blogs, newBlog)
	c.JSON(http.StatusCreated, newBlog)
}

// Updates the blogpost based on the id passed on
func updateBlogPost(c *gin.Context) {
	var newBlog blog
	id := c.Param("Id")
	db.Model(&blog{}).Where("id= ?", id).Updates(&newBlog)
	c.JSON(http.StatusOK, newBlog)
}

// delete the post based on the Id as primary key if not available will promt error
func deleteBlogPost(c *gin.Context) {
	id := c.Param("Id")
	db.Delete(&blog{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted"})
}

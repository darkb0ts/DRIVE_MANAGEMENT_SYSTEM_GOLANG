package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files from the "static" directory
	router.Static("/static", "./static")

	// Load HTML templates from the "templates" directory
	router.LoadHTMLGlob("templates/*")

	// Define a route to render the login page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{})
	})

	// Run the server on port 8080
	router.Run(":8080")
}

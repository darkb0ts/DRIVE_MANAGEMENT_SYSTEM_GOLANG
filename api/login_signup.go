package main

import (
	"fmt"
	"net/http"

	"/config/conection_data" // replace with the actual path if different

	"github.com/gin-gonic/gin"
)

// LoginSignupHandler handles the login/signup route
func LoginSignupHandler(c *gin.Context) {
	db := conection_data.ConnectToDB()

	// Your database operations here using "db"
	fmt.Print(db)

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func main() {
	// Start your Gin server
	router := gin.Default()

	// Define your routes
	router.GET("/login", LoginSignupHandler)

	// Run the server
	router.Run(":8080")
}

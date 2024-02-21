package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Drive struct represents the data model for a drive
type Drive struct {
    Name          string `json:"name"`
    VehicleNumber string `json:"vehicle_number"`
}

var db *sql.DB

func main() {
    // Connect to the database
    var err error
    db, err = sql.Open("mysql", "######################")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize Gin
    router := gin.Default()

    // Define API endpoints
    router.GET("/drives", getDrives)

    // Run the server
    router.Run(":8080")
}

func getDrives(c *gin.Context) {
    // Query the database to get drive names and vehicle numbers
    rows, err := db.Query("SELECT name, vehicle_number FROM drives")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    // Iterate over the rows and populate a slice of Drive structs
    var drives []Drive
    for rows.Next() {
        var drive Drive
        if err := rows.Scan(&drive.Name, &drive.VehicleNumber); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        drives = append(drives, drive)
    }

    // Check for errors during iteration
    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the list of drives as JSON
    c.JSON(http.StatusOK, drives)
}

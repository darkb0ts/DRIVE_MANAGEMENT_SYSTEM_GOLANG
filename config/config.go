package conection_data

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

func ConnectToDB() *pg.DB {
	// PostgreSQL connection options
	opts := &pg.Options{
		Addr:     "0.0.0.0:5432",
		User:     "postgres",
		Password: "mysecretpassword",
		Database: "postgres",
	}

	// Create a new PostgreSQL connection
	db := pg.Connect(opts)

	// Check if the connection is successful
	_, err := db.Exec("SELECT 1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to the database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to PostgreSQL!")

	return db
}

// func main() {
// 	// Connect to the PostgreSQL database
// 	db := ConnectToDB()

// 	// Close the database connection when done
// 	defer db.Close()
// }

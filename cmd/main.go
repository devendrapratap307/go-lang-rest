package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/restapi-go/cmd/api"
	"github.com/restapi-go/db"
)

func main() {
	// connect to the database
	dbConn, err := db.NewPostgresSQL(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled"))
	if err != nil {

	}
	if err := initDB(dbConn); err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// create api
	apiServer := api.NewAPIServer(":9080")
	if err := apiServer.Run(); err != nil {
		log.Fatal("Error starting API server: ", err)
	} else {
		fmt.Println("API server started successfully")
	}
}

func initDB(db *sql.DB) error {
	// Initialize the database connection
	fmt.Println("Database initialized")
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

package tictactoe 

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// Connect to Postgres DB
func OpenConnection() (*sql.DB, string) {
	//Get constants from env variables
	err := godotenv.Load()
	
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	host, ok := os.LookupEnv("DB_HOST")

	if !ok {
		log.Fatal("DB_HOST is not set in .env file")
	}

	user, ok := os.LookupEnv("DB_USER")

	if !ok {
		log.Fatal("DB_USER is not set in .env file")
	}

	port, ok := os.LookupEnv("DB_PORT")

	if !ok {
		log.Fatal("DB_PORT is not set in .env file")
	}

	password, ok := os.LookupEnv("DB_PASSWORD")

	if !ok {
		log.Fatal("DB_PASSWORD is not set in .env file")
	}

	dbName, ok := os.LookupEnv("DB_NAME")

	if !ok {
		log.Fatal("DB_NAME is not set in .env file")
	}

	// connect to database
	// create a connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
			host, port, user, password, dbName)

	// validate arguements
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	// open database connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}


}

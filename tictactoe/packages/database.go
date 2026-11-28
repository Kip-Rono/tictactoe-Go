package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDB() {
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

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	// configure connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	// open database connection
	err = db.Ping()

	if err != nil {
		return err
	}

	return db, nil
}

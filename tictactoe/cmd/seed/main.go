package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)


func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	defer db.Close()

	//impelement seeding logic here
	hash, _:= bcrypt.GenerateFromPassword([]byte(os.Getenv("SUPER_ADMIN_PASSWORD")), bcrypt.DefaultCost)

	_, err = db.Exec(`
		INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (email) DO NOTHING
	`, os.Getenv("SUPER_ADMIN_NAME"), os.Getenv("SUPER_ADMIN_EMAIL"), string(hash))

	if err != nil {
		log.Fatal("Error seeding the database: ", err)
	}

	log.Println("Database seeded successfully")
}
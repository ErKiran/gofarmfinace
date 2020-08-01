package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	dbType := os.Getenv("DB_TYPE")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	db, err := gorm.Open(dbType, fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		name,
		password,
	))

	if err != nil {
		panic(err)
	}

	db.Debug()
}

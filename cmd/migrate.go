package main

import (
	"farmfinance/models"
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

	db.Exec("create type transaction_type as enum('DEBIT','CREDIT'); ")

	db.Debug().AutoMigrate(&models.Crop{}, &models.Farm{}, &models.Ledger{}, &models.FarmType{})

	db.Model(&models.Ledger{}).AddForeignKey("crop_id", "crops(id)", "CASCADE", "CASCADE")

	db.Model(&models.Crop{}).AddForeignKey("farm_id", "farms(id)", "CASCADE", "CASCADE")

	db.Model(&models.Farm{}).AddForeignKey("farm_type_id", "farm_types(id)", "CASCADE", "CASCADE")

}

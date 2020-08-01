package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	//Let
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func Init() {
	dbType := os.Getenv("DB_TYPE")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open(dbType, fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		name,
		password,
	))

	if err != nil {
		log.Fatal(err, "Unable to connect to database")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

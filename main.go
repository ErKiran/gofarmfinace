package main

import (
	"farmfinance/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	models.Init()
}

func main() {

	fmt.Println("Hello World!!, something great is on the way!!")
}

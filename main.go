package main

import (
	"farmfinance/models"
	"farmfinance/routers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	mode := os.Getenv("RUN_MODE")
	port := os.Getenv("HTTP_PORT")

	port = fmt.Sprintf(":%v", port)

	fmt.Println("PORT", port)

	gin.SetMode(mode)
	routersInit := routers.InitRouter()
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           port,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

	fmt.Println("Hello World!!, something great is on the way!!")
}

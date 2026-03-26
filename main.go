package main

import (
	"api/connect"
	"api/products"
	"api/router"
	"log"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	Port := loadenv()
	DB,err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	products := products.NewModule(DB)

	router := router.RouterAPI(products.Controller,)
	router.Run(":" + Port)
	
}

func loadenv() string{
	err := godotenv.Load(".env")
	if err != nil{
		log.Println("Error Port")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	return PORT
}
package main

import (
	"api/connect"
	"api/products"
	"api/router"
	"log"
)

func main() {
	DB,err := connect.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	products := products.NewModule(DB)

	router := router.RouterAPI(products.Controller,)
	router.Run(":30606")
	
}
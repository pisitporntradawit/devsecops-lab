package router

import (
	"api/products"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func RouterAPI(
	ProductController *products.Controller,
) *gin.Engine{
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000","http://localhost:3001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	GroupProducts(r, ProductController)
	return r
}

func GroupProducts(r *gin.Engine, ProductController *products.Controller){
	productGroup := r.Group("/products")
	{
		productGroup.GET("", ProductController.GetProducts)
		productGroup.POST("",ProductController.CreateProduct)
	}
}
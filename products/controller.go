package products

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct{
	Service *Service
}

func NewController(service *Service) *Controller{
	return &Controller{
		Service :service,
	}
}

type Module struct{
	Controller *Controller
	Service *Service
	Repository *Repository
}

func NewModule(db *pgxpool.Pool) *Module{
	repo := NewRepository(db)
	service := NewService(repo)
	controller := NewController(service)

	return &Module{
		Controller: controller,
		Service: service,
		Repository: repo,
	}
}

func(ctrl *Controller) GetProducts(c *gin.Context){
	products, err := ctrl.Service.GetProducts(context.Background())
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func(ctrl *Controller) CreateProduct(c *gin.Context){
	var req ProductsModel
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	if err := ctrl.Service.CreateProduct(c.Request.Context(), &req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Name" : req.Name})
}
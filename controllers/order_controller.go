package controllers

import (
	"Final_Project/entity"
	"Final_Project/repositories"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderRepository repositories.OrderRepository
}

type OrderControllers interface {
	AllOrders(ctx *gin.Context)
	GetStatus(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

func NewOrderController(orderRepository *repositories.OrderRepository) *OrderController {
	return &OrderController{orderRepository: *orderRepository}
}

// AllOrders : Get all orders
func (c *OrderController) AllOrders(ctx *gin.Context) {
	result := c.orderRepository.GetAllOrders()
	ctx.JSON(200, entity.Response{
		Status: "success",
		Data:   result,
	})
}

func (c *OrderController) GetStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		log.Fatal(err)
	}

	orderStatus, err := c.orderRepository.GetStatus(id)
	ctx.JSON(200, entity.Response{
		Status:       "success",
		Data:         orderStatus,
		ErrorMessage: err.Error(),
	})
}

// CreateOrder Create new order
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var reqBody entity.Order
	err := ctx.Bind(&reqBody)
	if err != nil {
		log.Fatal(err)
	}

	c.orderRepository.CreateNewOrder(reqBody)
	ctx.JSON(200, entity.Response{
		Status: "Created",
		Data:   nil,
	})
}

// DeleteOrder by ID
func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	c.orderRepository.DeleteOrder(id)
	ctx.JSON(200, true)
}

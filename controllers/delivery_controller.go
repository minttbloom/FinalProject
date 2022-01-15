package controllers

import (
	"Final_Project/entity"
	"Final_Project/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeliveryController struct {
	orderRepository repositories.OrderRepository
}

type DeliveryControllers interface {
	CompleteDelivery(ctx *gin.Context)
}

func NewDeliveryController(orderRepository *repositories.OrderRepository) *DeliveryController {
	return &DeliveryController{orderRepository: *orderRepository}
}

// Change status delivery for order
func (c *DeliveryController) CompletionDelivery(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	c.orderRepository.ChangeStatus(id, entity.Complete)
	ctx.JSON(200, entity.Response{
		Status: "success",
		Data:   true,
	})
}

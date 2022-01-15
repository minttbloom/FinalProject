package controllers

import (
	"Final_Project/entity"
	"Final_Project/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KitchenController struct {
	orderRepository repositories.OrderRepository
}

type KitchenControllers interface {
	CompleteCook(ctx *gin.Context)
}

func NewKitchenController(orderRepository *repositories.OrderRepository) *KitchenController {
	return &KitchenController{
		orderRepository: *orderRepository,
	}
}

func (c *KitchenController) CompletionCook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	c.orderRepository.ChangeStatus(id, entity.OnRoad)
	ctx.JSON(200, entity.Response{
		Status: "success",
		Data:   true,
	})
}

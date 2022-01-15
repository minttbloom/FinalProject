package main

import (
	"Final_Project/controllers"
	"Final_Project/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := db.ConnectDB()
	defer db.Close()

	repositories := DefineRepositories(db)

	orderController := controllers.NewOrderController(&repositories.Order)
	deliveryController := controllers.NewDeliveryController(&repositories.Order)
	kitchenController := controllers.NewKitchenController(&repositories.Order)

	// Orders
	r.GET("order/all", orderController.AllOrders)
	r.GET("order/status", orderController.GetStatus)
	r.PUT("order/create", orderController.CreateOrder)
	r.DELETE("order/delete", orderController.DeleteOrder)

	// Kitchen
	r.POST("kitchen/completion", kitchenController.CompletionCook)

	// Delivery
	r.POST("delivery/completion", deliveryController.CompletionDelivery)

	r.Run()
}

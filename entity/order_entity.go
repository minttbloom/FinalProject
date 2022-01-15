package entity

import "time"

type OrderStatus int

const (
	Created  OrderStatus = 0
	Cooking  OrderStatus = 1
	OnRoad   OrderStatus = 2
	Complete OrderStatus = 3
)

type Order struct {
	Id              int    `bson:"id"`
	ClientName      string `bson:"client_name"`
	PizzaItemId     []int
	CreateTimeMs    *time.Time `bson:"create_time_ms"`
	CompleteTimeMs  *time.Time `bson:"complete_time_ms"`
	DeliveryAddress string     `bson:"delivery_address"`
	PhoneNumber     string     `bson:"phone_number"`
	Status          OrderStatus
}

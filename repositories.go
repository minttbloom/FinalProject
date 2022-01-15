package main

import (
	"Final_Project/repositories"
	"database/sql"
)

type Repositories struct {
	Order repositories.OrderRepository
}

func DefineRepositories(d *sql.DB) *Repositories {
	return &Repositories{
		Order: *repositories.NewOrderRepository(d),
	}
}

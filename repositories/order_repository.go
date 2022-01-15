package repositories

import (
	"Final_Project/entity"
	"context"
	"database/sql"
	"log"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) GetAllOrders() []entity.Order {
	var o []entity.Order
	ctx := context.Background()
	result, err := r.DB.QueryContext(ctx, "SELECT * FROM Orders")
	if err != nil {
		log.Fatal(err)
	}

	result.Scan(&o)
	return o
}

func (r *OrderRepository) GetStatus(id int) (*entity.OrderStatus, error) {
	ctx := context.Background()
	orderResult := r.DB.QueryRowContext(ctx, "SELECT * FROM Orders WHERE id=?", id)

	var o entity.Order
	err := orderResult.Scan(
		&o.Id,
		&o.ClientName,
		&o.CreateTimeMs,
		&o.CompleteTimeMs,
		&o.DeliveryAddress,
		&o.PhoneNumber,
	)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &o.Status, nil
}

func (r *OrderRepository) CreateNewOrder(e entity.Order) {
	ctx := context.Background()
	_, err := r.DB.QueryContext(ctx,
		"INSERT INTO Orders(client_name, pizza_id, create_time_ms, delivery_address, phone_number) VALUES (?, ?, ?, ?, ?)",
		e.ClientName, e.PizzaItemId, e.CreateTimeMs, e.DeliveryAddress, e.PhoneNumber)

	_, err = r.DB.QueryContext(ctx,
		"INSERT INTO")
	if err != nil {
		log.Fatal(err)
	}
}

func (r *OrderRepository) DeleteOrder(id int) {
	ctx := context.Background()
	_, err := r.DB.QueryContext(ctx, "DELETE FROM Orders WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *OrderRepository) ChangeStatus(orderId int, status entity.OrderStatus) {
	var err error

	ctx := context.Background()

	if status == entity.Complete {
		_, err = r.DB.ExecContext(ctx,
			"UPDATE OrderedPizza SET status=? WHERE order_id=?",
			status, orderId)

		_, err = r.DB.ExecContext(ctx,
			"UPDATE Orders SET complete_time_ms=CURRENT_TIMESTAMP WHERE id=?", orderId)

	} else {
		_, err = r.DB.ExecContext(ctx, "UPDATE OrderedPizza SET status=? WHERE order_id=?;", status, orderId)
	}

	if err != nil {
		log.Fatal(err)
	}
}

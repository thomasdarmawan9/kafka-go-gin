package order_domain

import (
	"fmt"
	"orders/db"
	"orders/utils/error_formats"
	"orders/utils/error_utils"
)

var OrderRepo orderDomain = &orderRepo{}

const (
	queryCreateOrder       = `INSERT INTO orders (foodId, userId, amount, total_price) VALUES ($1, $2, $3, $4) RETURNING id, foodId, userId, amount, total_price, created_at`
	queryGetOrdersByUserId = `SELECT id, foodId, userId, amount, total_price, status, created_at from orders WHERE userId = $1 AND status != 'PAID'`
	queryUpdateOrderStatus = `UPDATE orders SET status = $2 WHERE id = $1 RETURNING id, foodId, userId, amount, total_price, created_at`
	queryUpdateOrderAmount = `UPDATE orders SET amount = $2, total_price = $3 WHERE id = $1`
	queryGetOrderById      = `SELECT id, foodId, userId, amount, total_price, status, created_at from orders WHERE id = $1`
)

type orderDomain interface {
	CreateOrder(*Order) (*Order, error_utils.MessageErr)
	UpdateOrderStatus(*Order) error_utils.MessageErr
	GetOrdersByUserId(int32) ([]Order, error_utils.MessageErr)
	UpdateOrderAmount(*Order) (*Order, error_utils.MessageErr)
	GetOrderById(int32) (*Order, error_utils.MessageErr)
}

type orderRepo struct{}

func (o *orderRepo) CreateOrder(orderReq *Order) (*Order, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryCreateOrder, orderReq.FoodId, orderReq.UserId, orderReq.Amount, orderReq.TotalPrice)

	var order Order

	if err := row.Scan(&order.Id, &order.FoodId, &order.UserId, &order.Amount, &order.TotalPrice, &order.CreatedAt); err != nil {
		fmt.Println("Error:", err)
		return nil, error_formats.ParseError(err)
	}

	return &order, nil
}

func (o *orderRepo) UpdateOrderStatus(orderReq *Order) error_utils.MessageErr {
	db := db.GetDB()

	_, err := db.Exec(queryUpdateOrderStatus, orderReq.Id, orderReq.Status)

	if err != nil {
		return error_formats.ParseError(err)
	}

	return nil
}
func (o *orderRepo) GetOrdersByUserId(userId int32) ([]Order, error_utils.MessageErr) {
	db := db.GetDB()

	var orders []Order

	rows, err := db.Query(queryGetOrdersByUserId, userId)
	if err != nil {
		fmt.Println(err)
		return nil, error_formats.ParseError(err)
	}

	defer rows.Close()

	for rows.Next() {
		var order Order
		err = rows.Scan(&order.Id, &order.FoodId, &order.UserId, &order.Amount, &order.TotalPrice, &order.Status, &order.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return nil, error_formats.ParseError(err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (o *orderRepo) UpdateOrderAmount(orderReq *Order) (*Order, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryUpdateOrderAmount, orderReq.Id, orderReq.Amount, orderReq.TotalPrice)

	var order Order

	if err := row.Scan(&order.Id, &order.FoodId, &order.UserId, &order.Amount, &order.TotalPrice, &order.CreatedAt); err != nil {

		return nil, error_formats.ParseError(err)
	}

	return &order, nil
}

func (o *orderRepo) GetOrderById(orderId int32) (*Order, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryGetOrderById, orderId)

	var order Order

	if err := row.Scan(&order.Id, &order.FoodId, &order.UserId, &order.Amount, &order.TotalPrice, &order.Status, &order.CreatedAt); err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &order, nil
}

package order_service

import (
	"orders/domain/order_domain"
	"orders/utils/error_utils"
)

var OrderService orderServiceRepo = &orderService{}

type orderServiceRepo interface {
	CreateOrder(*order_domain.Order) (*order_domain.Order, error_utils.MessageErr)
	UpdateOrderStatus(*order_domain.Order) error_utils.MessageErr
	GetOrdersByUserId(int32) ([]order_domain.Order, error_utils.MessageErr)
	UpdateOrderAmount(*order_domain.Order) (*order_domain.Order, error_utils.MessageErr)
	GetOrderById(int32) (*order_domain.Order, error_utils.MessageErr)
}

type orderService struct{}

func (o *orderService) CreateOrder(orderReq *order_domain.Order) (*order_domain.Order, error_utils.MessageErr) {
	err := orderReq.Validate()

	if err != nil {
		return nil, err
	}

	order, err := order_domain.OrderRepo.CreateOrder(orderReq)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderService) UpdateOrderStatus(orderReq *order_domain.Order) error_utils.MessageErr {
	err := order_domain.OrderRepo.UpdateOrderStatus(orderReq)

	if err != nil {
		return err
	}

	return nil
}

func (o *orderService) GetOrdersByUserId(userId int32) ([]order_domain.Order, error_utils.MessageErr) {
	orders, err := order_domain.OrderRepo.GetOrdersByUserId(userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderService) UpdateOrderAmount(orderReq *order_domain.Order) (*order_domain.Order, error_utils.MessageErr) {
	order, err := order_domain.OrderRepo.UpdateOrderAmount(orderReq)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderService) GetOrderById(orderId int32) (*order_domain.Order, error_utils.MessageErr) {
	order, err := order_domain.OrderRepo.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}

	return order, nil
}

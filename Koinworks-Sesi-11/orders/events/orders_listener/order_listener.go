package orders_listener

import (
	"context"
	"encoding/json"
	"fmt"
	"orders/domain/food_domain"
	"orders/domain/order_domain"
	"orders/service/food_service"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	createOrder       = "CREATE-ORDER"
	createPayment     = "CREATE-PAYMENT"
	createFood        = "CREATE-FOOD"
	updateOrderStatus = "UPDATE-ORDER-STATUS"
	paymentFailed     = "PAYMENT-FAILED"
	paymentSucceeded  = "PAYMENT-SUCCEEDED"
)

const (
	orderService = "ORDERS-SERVICE"
)

var OrdersListener ordersListenerRepo = &ordersListener{}

type ordersListenerRepo interface {
	InitiliazeMainListener()
	CreateFood(string, []byte)
	PaymentFailed(string, []byte)
	PaymentSucceeded(string, []byte)
}

type ordersListener struct {
	kafka *kafka.Reader
}

func (o *ordersListener) InitiliazeMainListener() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	o.kafka = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   orderService,
		GroupID: "orders-group",
	})

	fmt.Println("Orders Service Start Listening")
	for {
		msg, err := o.kafka.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}

		key := string(msg.Key)
		value := msg.Value
		switch key {
		case createFood:
			go o.CreateFood(key, value)
		case paymentFailed:
			go o.PaymentFailed(key, value)
		case paymentSucceeded:
			go o.PaymentSucceeded(key, value)
		default:
			fmt.Println("key is not detected!")
		}

	}
}

func (o *ordersListener) CreateFood(key string, message []byte) {

	var food food_domain.Food
	err := json.Unmarshal(message, &food)

	fmt.Printf("%+v", food)

	if err != nil {
		fmt.Println("error unmarshalling json data:", err)
		return
	}

	errData := food_service.FoodService.CreateFood(&food)

	if errData != nil {
		fmt.Println("error create food:", errData.Message())
		return
	}

	fmt.Println("Food Created")
}

func (o *ordersListener) PaymentFailed(key string, message []byte) {
	var orders []order_domain.Order
	err := json.Unmarshal(message, &orders)

	if err != nil {
		fmt.Println("error unmarshalling json data:", err)
		return
	}

	for _, order := range orders {
		order.Status = "REJECTED"
		errData := order_domain.OrderRepo.UpdateOrderStatus(&order)
		if errData != nil {
			fmt.Println("error update order status:", errData.Message())
			return
		}
	}

	fmt.Println("order status rejected")
}

func (o *ordersListener) PaymentSucceeded(key string, message []byte) {
	var orders []order_domain.Order
	err := json.Unmarshal(message, &orders)

	if err != nil {
		fmt.Println("error unmarshalling json data:", err)
		return
	}

	for _, order := range orders {
		order.Status = "PAID"
		errData := order_domain.OrderRepo.UpdateOrderStatus(&order)
		if errData != nil {
			fmt.Println("error update order status:", errData.Message())
			return
		}

		var food food_domain.Food = food_domain.Food{
			Id:    order.FoodId,
			Stock: order.Amount,
		}
		errData = food_service.FoodService.ReduceFoodStock(&food)

		if errData != nil {
			fmt.Println("error update order status:", errData.Message())
			return
		}
	}

	fmt.Println("order status and food stock updated")
}

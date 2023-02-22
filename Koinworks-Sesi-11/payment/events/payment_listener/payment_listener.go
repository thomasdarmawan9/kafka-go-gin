package payment_listener

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"payment/events/payment_producer"

	"github.com/segmentio/kafka-go"
)

const (
	createPayment = "CREATE-PAYMENT"
)

const (
	listenerTopic = "PAYMENT-SERVICE"
)

var PaymentListener paymentListenerRepo = &paymentListener{}

type paymentListenerRepo interface {
	InitiliazeMainListener()
	CreatePayment(string, []byte)
}

type paymentListener struct {
	kafka *kafka.Reader
}

func (payment *paymentListener) InitiliazeMainListener() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	payment.kafka = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   listenerTopic,
		GroupID: "payment-group",
	})

	fmt.Println("Payment Listener: Start Listening")
	for {
		msg, err := payment.kafka.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}

		key := string(msg.Key)
		value := msg.Value

		switch key {
		case createPayment:
			payment.CreatePayment(key, value)
		default:
			fmt.Println("key is not detected!")
		}
	}
}

type Order struct {
	Id         int32   `json:"id"`
	FoodId     int32   `json:"food_id" valid:"required~food id is required"`
	UserId     int32   `json:"user_id" valid:"required~user id is required"`
	Amount     int8    `json:"amount" valid:"required~order amount is required"`
	TotalPrice float32 `json:"total_price"`
}

func (payment *paymentListener) CreatePayment(key string, value []byte) {
	var orders []Order

	err := json.Unmarshal(value, &orders)

	if err != nil {
		fmt.Println("error unmarshalling json data:", err)
		return
	}

	var totalPrice float32

	for _, order := range orders {
		totalPrice += order.TotalPrice
	}

	if totalPrice >= float32(200) {
		payment_producer.PaymentProducer.PaymentDone("PAYMENT-FAILED", value)
		return
	}

	payment_producer.PaymentProducer.PaymentDone("PAYMENT-SUCCEEDED", value)
}

package food_listener

import (
	"context"
	"encoding/json"
	"fmt"
	"foods/domain/food_domain"
	"foods/service/food_service"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	paymentSucceeded = "PAYMENT-SUCCEEDED"
)

const (
	foodsService = "FOODS-SERVICE"
)

var FoodListener foodsListenerRepo = &foodsListener{}

type foodsListenerRepo interface {
	InitiliazeMainListener()
	PaymentSucceeded(string, []byte)
}

type foodsListener struct {
	kafka *kafka.Reader
}

func (f *foodsListener) InitiliazeMainListener() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	f.kafka = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   foodsService,
		GroupID: "foods-group",
	})

	fmt.Println("Foods Service Start Listening")
	for {
		msg, err := f.kafka.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}

		key := string(msg.Key)
		value := msg.Value
		switch key {
		case paymentSucceeded:
			go f.PaymentSucceeded(key, value)
		default:
			fmt.Println("key is not detected!")
		}

	}
}

func (f *foodsListener) PaymentSucceeded(key string, message []byte) {
	var v []map[string]interface{}

	err := json.Unmarshal(message, &v)

	if err != nil {
		fmt.Println("error unmarshalling json data:", err)
		return
	}

	for _, food := range v {
		foodReq := &food_domain.Food{
			Id:    int32(food["id"].(float64)),
			Stock: int8(food["amount"].(float64)),
		}

		theErr := food_service.FoodService.ReduceFoodStock(foodReq)

		if theErr != nil {
			fmt.Println("error reducing food stock:", theErr.Message())
			return
		}
	}
	fmt.Println("Reducing Food Stock Succeeded")
}

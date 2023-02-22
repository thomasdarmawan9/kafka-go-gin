package orchestrator_producer

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

var OrchestratorProducer orchestratorProducerRepo = &orchestratorProducer{}

type orchestratorProducerRepo interface {
	SetUpProducer()
	CreateFood(string, []byte)
	CreatePayment(string, []byte)
	PaymentSucceeded(string, []byte)
	PaymentFailed(string, []byte)
}

type orchestratorProducer struct {
	kafka *kafka.Writer
}

func (orch *orchestratorProducer) SetUpProducer() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	l := log.New(os.Stdout, "orchestrator producer writer: ", 0)
	orch.kafka = &kafka.Writer{
		Addr:   kafka.TCP(brokerAddress),
		Logger: l,
	}
}

func (orch *orchestratorProducer) CreateFood(key string, message []byte) {
	err := orch.kafka.WriteMessages(context.Background(), kafka.Message{
		Topic: "ORDERS-SERVICE",
		Key:   []byte(key),
		Value: message,
	})
	if err != nil {
		fmt.Println("Cannot send message create food:", err)
	}
	fmt.Println("create food message has been sent")
}

func (orch *orchestratorProducer) CreatePayment(key string, message []byte) {

	err := orch.kafka.WriteMessages(context.Background(), kafka.Message{
		Topic: "PAYMENT-SERVICE",
		Key:   []byte(key),
		Value: message,
	})
	if err != nil {
		fmt.Println("Cannot send message create payment:", err)
	}

	fmt.Println("Create Payment Message has been sent")
}

func (orch *orchestratorProducer) PaymentFailed(key string, message []byte) {

	err := orch.kafka.WriteMessages(context.Background(), kafka.Message{
		Topic: "ORDERS-SERVICE",
		Key:   []byte(key),
		Value: message,
	})
	if err != nil {
		fmt.Println("Cannot send message PAYMENT-FAILED:", err)
	}
}

func (orch *orchestratorProducer) PaymentSucceeded(key string, message []byte) {

	err := orch.kafka.WriteMessages(context.Background(),
		kafka.Message{
			Topic: "ORDERS-SERVICE",
			Key:   []byte(key),
			Value: message,
		},
		kafka.Message{
			Topic: "FOODS-SERVICE",
			Key:   []byte(key),
			Value: message,
		},
	)
	if err != nil {
		fmt.Println("Cannot send message PAYMENT-SUCCEEDED:", err)
	}
}

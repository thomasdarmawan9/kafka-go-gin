package orders_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	orchestratorService = "ORCHESTRATOR-SERVICE"
)

var OrderProducer orderProducerRepo = &orderProducer{}

type orderProducerRepo interface {
	SetUpProducer()
	CreatePayment(string, interface{})
}

type orderProducer struct {
	kafka *kafka.Writer
}

func (order *orderProducer) SetUpProducer() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	order.kafka = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   orchestratorService,
	})
}
func (order *orderProducer) CreatePayment(key string, message interface{}) {
	value, _ := json.Marshal(message)

	err := order.kafka.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		fmt.Println("Cannot send message create payment:", err)
	}
	fmt.Println("create payment message has been sent")
}

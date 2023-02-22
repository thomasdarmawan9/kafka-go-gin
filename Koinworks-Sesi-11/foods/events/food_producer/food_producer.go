package food_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	orchestratorService = "ORCHESTRATOR-SERVICE"
)

var FoodProducer foodProducerRepo = &foodProducer{}

type foodProducerRepo interface {
	SetUpProducer()
	CreateFood(string, interface{})
}

type foodProducer struct {
	kafka *kafka.Writer
}

func (orch *foodProducer) SetUpProducer() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	l := log.New(os.Stdout, "food producer writer: ", 0)
	orch.kafka = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   orchestratorService,
		Logger:  l,
	})
}

func (orch *foodProducer) CreateFood(key string, message interface{}) {
	value, _ := json.Marshal(message)

	err := orch.kafka.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		fmt.Println("Cannot send message create food:", err)
	}
	fmt.Println("create food message has been sent")
}

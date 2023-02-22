package payment_producer

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	orchestratorService = "ORCHESTRATOR-SERVICE"
)

var PaymentProducer paymentProducerRepo = &paymentProducer{}

type paymentProducerRepo interface {
	SetUpProducer()
	PaymentDone(string, []byte)
}

type paymentProducer struct {
	kafka *kafka.Writer
}

func (p *paymentProducer) SetUpProducer() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	p.kafka = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   orchestratorService,
	})
}

func (p *paymentProducer) PaymentDone(key string, message []byte) {
	err := p.kafka.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: message,
	})

	if err != nil {
		fmt.Printf("Cannot send %s message \n", key)
	}

	fmt.Println(key, "message has been successsfully sent")
}

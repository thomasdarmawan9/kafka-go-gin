package orchestrator_listener

import (
	"context"
	"fmt"
	"orchestrator/events/orchestrator_producer"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	createOrder      = "CREATE-ORDER"
	createPayment    = "CREATE-PAYMENT"
	createFood       = "CREATE-FOOD"
	paymentFailed    = "PAYMENT-FAILED"
	paymentSucceeded = "PAYMENT-SUCCEEDED"
)

const (
	orchestratorService = "ORCHESTRATOR-SERVICE"
)

var OrchestatorListener orcherstratorListenerRepo = &orchestatorListener{}

type orcherstratorListenerRepo interface {
	InitiliazeMainListener()
}

type orchestatorListener struct {
	kafka *kafka.Reader
}

func (orch *orchestatorListener) InitiliazeMainListener() {
	brokerAddress := os.Getenv("BROKER_ADDRESS")
	orch.kafka = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   orchestratorService,
		GroupID: "orchestator-group",
	})

	fmt.Println("Orchestator Listener: Start Listening")
	for {
		msg, err := orch.kafka.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("could not read message " + err.Error())
		}

		key := string(msg.Key)
		value := msg.Value
		switch key {
		case createPayment:
			orchestrator_producer.OrchestratorProducer.CreatePayment(key, value)
		case createFood:
			orchestrator_producer.OrchestratorProducer.CreateFood(key, value)
		case paymentFailed:
			orchestrator_producer.OrchestratorProducer.PaymentFailed(key, value)
		case paymentSucceeded:
			orchestrator_producer.OrchestratorProducer.PaymentSucceeded(key, value)
		default:
			fmt.Println("key is not detected! ==>", key)
		}

	}
}

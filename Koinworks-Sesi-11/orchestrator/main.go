package main

import (
	"orchestrator/events/orchestrator_listener"
	"orchestrator/events/orchestrator_producer"
)

func main() {
	orchestrator_producer.OrchestratorProducer.SetUpProducer()
	orchestrator_listener.OrchestatorListener.InitiliazeMainListener()
}

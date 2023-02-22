package main

import (
	"payment/events/payment_listener"
	"payment/events/payment_producer"
)

func main() {
	payment_producer.PaymentProducer.SetUpProducer()
	payment_listener.PaymentListener.InitiliazeMainListener()
}

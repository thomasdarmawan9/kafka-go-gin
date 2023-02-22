package main

import (
	"orders/events/orders_listener"
	"orders/router"
)

func main() {
	go orders_listener.OrdersListener.InitiliazeMainListener()
	router.StartRouter()
}

package router

import (
	"orders/controller/order_controller"
	"orders/db"
	"orders/events/orders_producer"
	"orders/middleware"

	"github.com/gin-gonic/gin"
)

var PORT = ":8082"

func init() {
	orders_producer.OrderProducer.SetUpProducer()
	db.InitializeDB()
}

func StartRouter() {
	route := gin.Default()

	orderRoute := route.Group("/orders")
	{
		orderRoute.Use(middleware.Authentication())
		//Create New Order
		orderRoute.POST("/", order_controller.CreateOrder)
		//Get All Orders Owned By Current Logged In User That Hasn't Been Paid Yet
		orderRoute.GET("/", order_controller.GetOrdersByUserId)
		//Update Amount Of One Order Data That Hasn't Been Paid Yet
		orderRoute.PUT("/:orderId", middleware.OrderAuthorization(), order_controller.UpdateOrderAmount)
		// Checkout Orders
		orderRoute.POST("/pay-orders", order_controller.PayOrders)
	}

	route.Run(PORT)

}

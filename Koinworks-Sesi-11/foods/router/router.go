package router

import (
	"foods/controller/food_controller"
	"foods/db"
	"foods/events/food_listener"
	"foods/events/food_producer"
	"foods/middleware"

	"github.com/gin-gonic/gin"
)

var PORT = ":8081"

func init() {
	db.InitializeDB()
	food_producer.FoodProducer.SetUpProducer()
}

func StartRouter() {
	route := gin.Default()

	foodRoute := route.Group("/foods")
	{
		foodRoute.Use(middleware.Authentication())
		foodRoute.POST("/", middleware.AdminAuthorization(), food_controller.CreateFood)
	}

	go food_listener.FoodListener.InitiliazeMainListener()
	route.Run(PORT)
}

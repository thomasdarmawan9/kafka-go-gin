package router

import (
	"users/controllers/user_controller"
	"users/db"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func init() {
	db.InitializeDB()
}

func StartRouter() {
	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", user_controller.UserRegister)
		userRoute.POST("/login", user_controller.UserLogin)
		userRoute.POST("/admin-register", user_controller.AdminRegister)
	}

	route.Run(PORT)
}

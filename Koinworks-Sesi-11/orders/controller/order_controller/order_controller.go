package order_controller

import (
	"fmt"
	"net/http"
	"orders/domain/order_domain"
	"orders/events/orders_producer"
	"orders/service/food_service"
	"orders/service/order_service"
	"orders/utils/error_utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateOrder(c *gin.Context) {
	var orderReq order_domain.Order

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	food, err := food_service.FoodService.GetFoodById(orderReq.FoodId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if food.Stock < orderReq.Amount {
		theErr := error_utils.NewBadRequest("insufficient food stock")
		c.JSON(theErr.Status(), theErr)
		return
	}

	userData := c.MustGet("userData").(jwt.MapClaims)

	orderReq.UserId = int32(userData["id"].(float64))

	orderReq.TotalPrice = food.Price * float32(orderReq.Amount)

	order, err := order_service.OrderService.CreateOrder(&orderReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, order)
}

func PayOrders(c *gin.Context) {
	var orderReq []order_domain.Order

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		fmt.Println(orderReq)
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	for _, order := range orderReq {
		if order.Id == 0 {
			theErr := error_utils.NewBadRequest("order id required")
			c.JSON(theErr.Status(), theErr)
			return
		}

		errData := order.Validate()

		if errData != nil {
			c.JSON(errData.Status(), errData)
			return
		}

		food, err := food_service.FoodService.GetFoodById(order.FoodId)
		if err != nil {
			c.JSON(err.Status(), err)
			return
		}

		if food.Stock < order.Amount {
			theErr := error_utils.NewBadRequest("food stock insufficient")
			c.JSON(theErr.Status(), theErr)
			return
		}
	}

	go orders_producer.OrderProducer.CreatePayment("CREATE-PAYMENT", orderReq)

	c.JSON(http.StatusOK, gin.H{
		"msg": "waiting for the payment process",
	})
}

func UpdateOrderAmount(c *gin.Context) {
	var orderReq order_domain.Order

	if err := c.ShouldBindJSON(&orderReq); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	orderId, err := orderReq.GetorderIdParam(c)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	food, err := food_service.FoodService.GetFoodById(orderReq.FoodId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if food.Stock < orderReq.Amount {
		theErr := error_utils.NewBadRequest("insufficient food stock")
		c.JSON(theErr.Status(), theErr)
		return
	}

	orderReq.Id = *orderId
	orderReq.TotalPrice = float32(orderReq.Amount) * food.Price

	order, err := order_domain.OrderRepo.UpdateOrderAmount(&orderReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, order)
}

func GetOrdersByUserId(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := int32(userData["id"].(float64))

	orders, err := order_domain.OrderRepo.GetOrdersByUserId(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, orders)
}

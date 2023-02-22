package order_domain

import (
	"orders/utils/error_utils"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Order struct {
	Id         int32     `json:"id"`
	FoodId     int32     `json:"food_id" valid:"required~food id is required"`
	UserId     int32     `json:"user_id" valid:"required~user id is required"`
	Amount     int8      `json:"amount" valid:"required~order amount is required"`
	TotalPrice float32   `json:"total_price" valid:"required~total price is required"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func (order *Order) Validate() error_utils.MessageErr {
	_, err := govalidator.ValidateStruct(order)

	if err != nil {
		return error_utils.NewBadRequest(err.Error())
	}

	return nil
}

func (order *Order) GetorderIdParam(c *gin.Context) (*int32, error_utils.MessageErr) {
	idParam := c.Param("orderId")
	orderId, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, error_utils.NewBadRequest("invalid order id params")
	}

	var result int32 = int32(orderId)

	return &result, nil
}

package food_domain

import (
	"foods/utils/error_utils"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Food struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name" valid:"required~food name is required"`
	Price     float32   `json:"price" valid:"required~food price is required"`
	Stock     int8      `json:"stock" valid:"required~food stock is required"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *Food) Validate() error_utils.MessageErr {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return error_utils.NewBadRequest(err.Error())
	}

	return nil
}

func (f *Food) GetfoodIdParam(c *gin.Context) (*int32, error_utils.MessageErr) {
	idParam := c.Param("foodId")
	foodId, err := strconv.Atoi(idParam)
	if err != nil {
		return nil, error_utils.NewBadRequest("invalid food id params")
	}

	var result int32 = int32(foodId)

	return &result, nil
}

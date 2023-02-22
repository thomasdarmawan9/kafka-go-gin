package food_controller

import (
	"fmt"
	"foods/domain/food_domain"
	"foods/events/food_producer"
	"foods/service/food_service"
	"foods/utils/error_utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {
	var foodReq food_domain.Food

	if err := c.ShouldBindJSON(&foodReq); err != nil {
		theErrr := error_utils.NewUnprocessibleEntityError("invalid json body")
		fmt.Println(err)
		c.JSON(theErrr.Status(), theErrr)
		return
	}

	res, err := food_service.FoodService.CreateFood(&foodReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	go food_producer.FoodProducer.CreateFood("CREATE-FOOD", res)
	c.JSON(http.StatusOK, res)
}

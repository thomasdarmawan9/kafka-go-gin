package user_controller

import (
	"net/http"
	"users/doc_datas"
	"users/domain/user_domain"
	"users/service/user_service"
	"users/utils/error_utils"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userReq user_domain.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErrr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErrr.Status(), theErrr)
		return
	}

	userReq.Role = "CUSTOMER"
	userData, err := user_service.UserService.CreateUser(&userReq)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := doc_datas.UserRegisterResponse{
		Id:        userData.Id,
		Email:     userData.Email,
		Address:   userData.Address,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

func AdminRegister(c *gin.Context) {
	var userReq user_domain.User
	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErrr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErrr.Status(), theErrr)
		return
	}

	userReq.Role = "ADMIN"
	userData, err := user_service.UserService.CreateUser(&userReq)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := doc_datas.UserRegisterResponse{
		Id:        userData.Id,
		Email:     userData.Email,
		Address:   userData.Address,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(http.StatusCreated, res)
}

func UserLogin(c *gin.Context) {
	var userReq user_domain.User

	if err := c.ShouldBindJSON(&userReq); err != nil {
		theErrr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErrr.Status(), theErrr)
		return
	}

	token, err := user_service.UserService.UserLogin(&userReq)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

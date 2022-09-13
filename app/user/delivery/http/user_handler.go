package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/t0239184/GoChat/app/domain"
	"github.com/t0239184/GoChat/app/tool"
)

type UserHandler struct {
	UserUsecase domain.IUserUsecase
}

type CreateUserRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func NewUserHandler(e *gin.Engine, userUsecase domain.IUserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}
	e.POST("/api/v1/user", handler.CreateUser)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	request := &CreateUserRequest{}
	tool.Logger.Info("CreateUser start.")
	if err := c.Bind(request); err != nil {
		fmt.Println(err)
		return
	}
	user := &domain.User{
		Account:  request.Account,
		Password: request.Password,
		Status:   "0",
	}

	id, err := u.UserUsecase.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
		"data": gin.H{
			"id": id,
		},
	})
}

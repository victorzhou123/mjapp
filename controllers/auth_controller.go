package controllers

import (
	"net/http"

	"mjbackend/models"
	"mjbackend/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{
		userService: services.NewUserService(),
	}
}

// 用户注册
func (ctrl *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	user, err := ctrl.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	// 返回用户信息（不包含密码）
	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("注册成功", gin.H{"user": userResponse}))
}

// 用户登录
func (ctrl *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	loginResponse, err := ctrl.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("登录成功", loginResponse))
}
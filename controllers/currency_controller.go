package controllers

import (
	"net/http"
	"strings"

	"mjbackend/models"
	"mjbackend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrencyController struct {
	currencyService *services.CurrencyService
}

func NewCurrencyController(currencyService *services.CurrencyService) *CurrencyController {
	return &CurrencyController{
		currencyService: currencyService,
	}
}

// GetBalance 查询算力余额
func (ctrl *CurrencyController) GetBalance(c *gin.Context) {
	// 从JWT中获取用户ID
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("未授权，请先登录"))
		return
	}
	
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse("用户ID格式错误"))
		return
	}
	
	// 查询余额
	balance, err := ctrl.currencyService.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse("查询余额失败: "+err.Error()))
		return
	}
	
	c.JSON(http.StatusOK, models.SuccessWithMessage("查询成功", balance))
}

// DeductBalance 扣减算力
func (ctrl *CurrencyController) DeductBalance(c *gin.Context) {
	// 从JWT中获取用户ID
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("未授权，请先登录"))
		return
	}
	
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse("用户ID格式错误"))
		return
	}
	
	// 绑定请求参数
	var request models.DeductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("请求参数错误: "+err.Error()))
		return
	}
	
	// 扣减算力
	result, err := ctrl.currencyService.DeductBalance(userID, &request)
	if err != nil {
		// 检查是否是余额不足错误
		if strings.Contains(err.Error(), "算力余额不足") {
			// 解析当前余额和需要的金额
			balanceResp, _ := ctrl.currencyService.GetBalance(userID)
			currentBalance := 0
			if balanceResp != nil {
				currentBalance = balanceResp.Balance
			}
			
			errorData := models.InsufficientBalanceError{
				CurrentBalance: currentBalance,
				RequiredAmount: request.Amount,
			}
			c.JSON(http.StatusBadRequest, models.ErrorResponseWithData(400, "算力余额不足", errorData))
			return
		}
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse("扣减算力失败: "+err.Error()))
		return
	}
	
	c.JSON(http.StatusOK, models.SuccessWithMessage("扣减成功", result))
}

// RechargeBalance 充值算力
func (ctrl *CurrencyController) RechargeBalance(c *gin.Context) {
	// 从JWT中获取用户ID
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("未授权，请先登录"))
		return
	}
	
	userID, ok := userIDInterface.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse("用户ID格式错误"))
		return
	}
	
	// 绑定请求参数
	var request models.RechargeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("请求参数错误: "+err.Error()))
		return
	}
	
	// 充值算力
	result, err := ctrl.currencyService.RechargeBalance(userID, &request)
	if err != nil {
		if strings.Contains(err.Error(), "交易ID已存在") {
			c.JSON(http.StatusBadRequest, models.BadRequestResponse("交易ID已存在，请勿重复充值"))
			return
		}
		if strings.Contains(err.Error(), "支付凭证验证失败") {
			c.JSON(http.StatusBadRequest, models.BadRequestResponse("支付凭证验证失败"))
			return
		}
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse("充值算力失败: "+err.Error()))
		return
	}
	
	c.JSON(http.StatusOK, models.SuccessWithMessage("充值成功", result))
}
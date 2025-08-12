package controllers

import (
	"net/http"
	"strconv"

	"mjbackend/middleware"
	"mjbackend/models"
	"mjbackend/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MemoController struct {
	memoService *services.MemoService
}

func NewMemoController() *MemoController {
	return &MemoController{
		memoService: services.NewMemoService(),
	}
}

// 创建备忘录
func (ctrl *MemoController) CreateMemo(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("用户未认证"))
		return
	}

	var req models.CreateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	memo, err := ctrl.memoService.CreateMemo(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("创建成功", memo))
}

// 获取备忘录列表
func (ctrl *MemoController) GetMemoList(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("用户未认证"))
		return
	}

	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	keyword := c.Query("keyword")

	// 参数验证
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	memoList, err := ctrl.memoService.GetMemoList(userID, page, limit, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.InternalServerErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("获取成功", memoList))
}

// 获取备忘录详情
func (ctrl *MemoController) GetMemoByID(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("用户未认证"))
		return
	}

	// 获取备忘录ID
	memoIDStr := c.Param("id")
	memoID, err := primitive.ObjectIDFromHex(memoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("无效的备忘录ID"))
		return
	}

	memo, err := ctrl.memoService.GetMemoByID(userID, memoID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NotFoundResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("获取成功", memo))
}

// 更新备忘录
func (ctrl *MemoController) UpdateMemo(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("用户未认证"))
		return
	}

	// 获取备忘录ID
	memoIDStr := c.Param("id")
	memoID, err := primitive.ObjectIDFromHex(memoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("无效的备忘录ID"))
		return
	}

	var req models.UpdateMemoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse(err.Error()))
		return
	}

	memo, err := ctrl.memoService.UpdateMemo(userID, memoID, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NotFoundResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("更新成功", memo))
}

// 删除备忘录
func (ctrl *MemoController) DeleteMemo(c *gin.Context) {
	userID, exists := middleware.GetUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse("用户未认证"))
		return
	}

	// 获取备忘录ID
	memoIDStr := c.Param("id")
	memoID, err := primitive.ObjectIDFromHex(memoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequestResponse("无效的备忘录ID"))
		return
	}

	err = ctrl.memoService.DeleteMemo(userID, memoID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.NotFoundResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessWithMessage("删除成功", nil))
}
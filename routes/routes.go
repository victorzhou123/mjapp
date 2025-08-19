package routes

import (
	"mjbackend/controllers"
	"mjbackend/middleware"
	"mjbackend/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 添加中间件
	r.Use(middleware.CORSMiddleware())

	// 创建服务实例
	currencyService := services.NewCurrencyService()

	// 创建控制器实例
	authController := controllers.NewAuthController()
	memoController := controllers.NewMemoController()
	currencyController := controllers.NewCurrencyController(currencyService)

	// API路由组
	api := r.Group("/api")
	{
		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		// 备忘录路由（需要认证）
		memos := api.Group("/memos")
		memos.Use(middleware.AuthMiddleware())
		{
			memos.GET("", memoController.GetMemoList)
			memos.POST("", memoController.CreateMemo)
			memos.GET("/:id", memoController.GetMemoByID)
			memos.PUT("/:id", memoController.UpdateMemo)
			memos.DELETE("/:id", memoController.DeleteMemo)
		}

		// 算力管理路由（需要认证）
		currency := api.Group("/currency")
		currency.Use(middleware.AuthMiddleware())
		{
			currency.GET("/balance", currencyController.GetBalance)
			currency.POST("/deduct", currencyController.DeductBalance)
			currency.POST("/recharge", currencyController.RechargeBalance)
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "备忘录后端服务运行正常",
			"version": "1.0.0",
		})
	})
}
package routes

import (
	"mjbackend/controllers"
	"mjbackend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 添加中间件
	r.Use(middleware.CORSMiddleware())

	// 创建控制器实例
	authController := controllers.NewAuthController()
	memoController := controllers.NewMemoController()

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
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "服务运行正常",
		})
	})
}
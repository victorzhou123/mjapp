package main

import (
	"log"

	"mjbackend/config"
	"mjbackend/database"
	"mjbackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 设置Gin模式
	gin.SetMode(config.AppConfig.GinMode)

	// 连接数据库
	database.ConnectMongoDB()

	// 创建Gin引擎
	r := gin.Default()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	log.Printf("服务器启动在端口 %s", config.AppConfig.Port)
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
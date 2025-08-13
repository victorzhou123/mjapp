package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"mjbackend/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongoDB() {
	// 创建MongoDB客户端选项
	fmt.Printf("config.AppConfig.MongoURI: %v\n", config.AppConfig.MongoURI)
	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoURI)

	// 连接到MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// 测试连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Connected to MongoDB successfully")

	// 设置数据库
	DB = client.Database(config.AppConfig.MongoDatabase)
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
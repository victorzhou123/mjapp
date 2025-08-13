package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port            string
	GinMode         string
	MongoURI        string
	MongoDatabase   string
	JWTSecret       string
	JWTExpiresHours int
	BcryptCost      int
}

var AppConfig *Config

func LoadConfig() {
	// 打印当前目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory:", err)
	}
	log.Printf("Current directory: %s", dir)

	// 解析JWT过期时间
	jwtExpires, err := strconv.Atoi(getEnv("JWT_EXPIRES_HOURS", "24"))
	if err != nil {
		jwtExpires = 24
	}

	// 解析bcrypt成本
	bcryptCost, err := strconv.Atoi(getEnv("BCRYPT_COST", "12"))
	if err != nil {
		bcryptCost = 12
	}

	AppConfig = &Config{
		Port:            getEnv("PORT", "8080"),
		GinMode:         getEnv("GIN_MODE", "debug"),
		MongoURI:        getEnv("MONGO_URI", "mongodb://admin:Password@1@192.168.22.113:30017"),
		MongoDatabase:   getEnv("MONGO_DATABASE", "memo_app"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key-here"),
		JWTExpiresHours: jwtExpires,
		BcryptCost:      bcryptCost,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

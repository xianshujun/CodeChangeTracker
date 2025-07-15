package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config 应用配置结构
type Config struct {
	Environment  string
	Port         string
	DatabaseURL  string
	RedisURL     string
	LogLevel     string
	JWTSecret    string
	OpenAIAPIKey string
}

// Load 加载配置
func Load() *Config {
	// 尝试加载.env文件
	_ = godotenv.Load()

	return &Config{
		Environment:  getEnv("ENVIRONMENT", "development"),
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://localhost/codechange_tracker?sslmode=disable"),
		RedisURL:     getEnv("REDIS_URL", "redis://localhost:6379"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvAsBool 获取布尔类型环境变量
func GetEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

// GetEnvAsInt 获取整数类型环境变量
func GetEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
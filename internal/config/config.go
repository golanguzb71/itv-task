package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	AppPort    int
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
	JWTSecret  string
	JWTExpiry  time.Duration
	AdminUser  string
	AdminPass  string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	appEnv := getEnv("APP_ENV", "development")
	appPort, _ := strconv.Atoi(getEnv("APP_PORT", "8080"))
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	dbName := getEnv("DB_NAME", "movies_db")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")
	jwtSecret := getEnv("JWT_SECRET", "default_secret_key")
	jwtExpiry, _ := time.ParseDuration(getEnv("JWT_EXPIRATION", "24h"))
	adminUser := getEnv("ADMIN_USERNAME", "admin")
	adminPass := getEnv("ADMIN_PASSWORD", "adminpassword")

	return &Config{
		AppEnv:     appEnv,
		AppPort:    appPort,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBSSLMode:  dbSSLMode,
		JWTSecret:  jwtSecret,
		JWTExpiry:  jwtExpiry,
		AdminUser:  adminUser,
		AdminPass:  adminPass,
	}, nil
}

func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}

func (c *Config) GetAppAddress() string {
	return fmt.Sprintf(":%d", c.AppPort)
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

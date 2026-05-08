package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

type Config struct {
	Db     DbConfig
	Auth   AuthConfig
	Server ServerConfig
	Logger LoggerConfig
	Redis  RedisConfig
}
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}
type AuthConfig struct {
	SecretKey     string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
	Issuer        string
}
type ServerConfig struct {
	ApiPort    string
	ApiHost    string
	StorageDir string
}

type LoggerConfig struct {
	Level  int
	Format string
}

type DbConfig struct {
	Dsn    string
	Logger int
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("не найден .env файл: %v \n", err)
		return
	}
	log.Println(".env загружен")
}

func LoadConfig() *Config {
	return &Config{
		Db: DbConfig{
			Dsn:    getString("DSN", ""),
			Logger: getInt("DB_LOGGER", int(logger.Info)),
		},
		Auth: AuthConfig{
			SecretKey:     getString("SECRET_JWT", "SECRET_JWT"),
			AccessExpiry:  time.Duration(int(time.Hour) * getInt("ACCESS_AXPIRY_JWT", 24)),
			RefreshExpiry: time.Duration(int(time.Hour) * getInt("REFRESH_AXPIRY_JWT", 24*7)),
			Issuer:        getString("ISSUER_JWT", "home-services"),
		},
		Server: ServerConfig{
			ApiPort:    getString("API_PORT", ":8080"),
			ApiHost:    getString("API_HOST", "localhost"),
			StorageDir: getString("STORAGE_DIR", "/mnt"),
		},
		Logger: LoggerConfig{
			Level:  getInt("LOG_LEVEL", 0),
			Format: getString("LOG_FORMAT", "json"),
		},
		Redis: RedisConfig{
			Host:     getString("REDIS_HOST", "localhost"),
			Port:     getInt("REDIS_PORT", 6379),
			Password: getString("REDIS_PASSWORD", ""),
			DB:       getInt("REDIS_DB", 0),
		},
	}
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	result, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return result
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	result, err := strconv.ParseBool(value)

	if err != nil {
		return defaultValue
	}

	return result
}

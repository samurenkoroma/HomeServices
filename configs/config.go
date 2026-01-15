package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Db     DbConfig
	Auth   AuthConfig
	Server ServerConfig
	Logger LoggerConfig
}

type AuthConfig struct {
	AccessSecret  string
	RefreshSecret string
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
	Dsn string
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
			Dsn: getString("DSN", ""),
		},
		Auth: AuthConfig{
			AccessSecret:  os.Getenv("ACCESS_SECRET_JWT"),
			RefreshSecret: os.Getenv("REFRESH_SECRET_JWT"),
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

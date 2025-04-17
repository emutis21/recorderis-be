package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	JWTSecret string

	Port string
	Env  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: No .env file found: %v", err)
	}

	config := &Config{
		JWTSecret: getEnvOrDefault("JWT_SECRET", "mi-super-secreto-de-desarrollo-123"),

		Port: getEnvOrDefault("PORT", "4000"),
		Env:  getEnvOrDefault("ENV", "development"),

		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrDefault("DB_USER", "recorderis_user"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "recorderis_pass"),
		DBName:     getEnvOrDefault("DB_NAME", "recorderis_db"),
	}

	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func NewDBConnection() (*sql.DB, error) {
	config := DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "recorderis_user",
		Password: "recorderis_pass",
		DBName:   "recorderis_db",
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	return db, nil
}

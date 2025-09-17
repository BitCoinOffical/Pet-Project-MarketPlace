package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	//postgreSQL
	PostgresUser    string
	PostgresPass    string
	PostgresDB      string
	PostgresHost    string
	PostgresPort    string
	PostgresSslmode string
	//redis
	RedisHost string
	RedisPort string
	//mailSetting
	Mail         string
	MailPassword string
	//api
	ApiPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	cfg := &Config{
		//postgreSQL
		PostgresUser: getEnv("POSTGRES_USER"),
		PostgresPass: getEnv("POSTGRES_PASSWORD"),
		PostgresDB:   getEnv("POSTGRES_DB"),
		PostgresHost: getEnv("POSTGRES_HOST"),
		PostgresPort: getEnv("POSTGRES_PORT"),
		//redis
		RedisHost: getEnv("REDIS_HOST"),
		RedisPort: getEnv("REDIS_PORT"),
		//mailSetting
		Mail:         getEnv("MAIL"),
		MailPassword: getEnv("MAIL_PASSWORD"),
		//api
		ApiPort: getEnv("API_PORT"),
	}
	return cfg
}
func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return value
}

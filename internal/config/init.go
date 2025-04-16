package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func loadEnvVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

type Config struct {
	POSTGRES_HOST      string
	POSTGRES_PORT      string
	POSTGRES_USER      string
	POSTGRES_PASSWORD  string
	POSTGRES_NAME      string
	POSTGRES_URL       string
	ACCESS_KEY         string
	ACCESS_TIME        time.Duration
	REFRESH_KEY        string
	REFRESH_TIME       time.Duration
	APP_VERSION        string
	PORT               string
	GIN_MODE           string
	LOGGER_FOLDER_PATH string
	LOGGER_FILENAME    string
	MOCK_REPOSITORY    string
	MOCK_DB            string
	MOCK_LOGGER        string
	MOCK_HANDLER       string
	MOCK_SERVICE       string
}

var ENV Config

func Init() *Config {
	godotenv.Load()

	ENV.PORT = loadEnvVariable("PORT")

	ENV.POSTGRES_HOST = loadEnvVariable("POSTGRES_HOST")
	ENV.POSTGRES_PORT = loadEnvVariable("POSTGRES_PORT")
	ENV.POSTGRES_USER = loadEnvVariable("POSTGRES_USER")
	ENV.POSTGRES_PASSWORD = loadEnvVariable("POSTGRES_PASSWORD")
	ENV.POSTGRES_NAME = loadEnvVariable("POSTGRES_NAME")

	ENV.GIN_MODE = loadEnvVariable("GIN_MODE")

	ENV.LOGGER_FOLDER_PATH = loadEnvVariable("LOGGER_FOLDER_PATH")
	ENV.LOGGER_FILENAME = loadEnvVariable("LOGGER_FILENAME")

	ENV.ACCESS_KEY = loadEnvVariable("ACCESS_KEY")
	AT, _ := time.ParseDuration(loadEnvVariable(("ACCESS_TIME")))
	ENV.ACCESS_TIME = AT
	ENV.REFRESH_KEY = loadEnvVariable("REFRESH_KEY")
	RT, _ := time.ParseDuration(loadEnvVariable(("REFRESH_TIME")))
	ENV.REFRESH_TIME = RT

	ENV.APP_VERSION = loadEnvVariable("APP_VERSION")

	ENV.MOCK_REPOSITORY = loadEnvVariable("MOCK_REPOSITORY")
	ENV.MOCK_DB = loadEnvVariable("MOCK_DB")
	ENV.MOCK_LOGGER = loadEnvVariable("MOCK_LOGGER")
	ENV.MOCK_HANDLER = loadEnvVariable("MOCK_HANDLER")
	ENV.MOCK_SERVICE = loadEnvVariable("MOCK_SERVICE")
	return &ENV
}

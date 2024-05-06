package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_KEY                 string `env:"API_KEY" env-required:"false"`
	TRANSCRIPTION_ENDPOINT  string `env:"TRANSCRIPTION_ENDPOINT" env-required:"false"`
	MODEL                   string `env:"MODEL" env-required:"false"`
	SERVER_URL              string `env:"SERVER_URL" env-required:"false"`
	RESPONSE_FORMAT         string `env:"RESPONSE_FORMAT" env-required:"false"`
	TIMESTAMP_GRANULARITIES string `env:"TIMESTAMP_GRANULARITIES" env-required:"false"`
	COMPLETION_ENDPOINT		string `env:"COMPLETION_ENDPOINT" env-required:"false"`
	COMPLETION_MODEL		string `env:"COMPLETION_MODEL" env-required:"false"`
	COMPLETION_MESSAGE		string `env:"COMPLETION_MESSAGE" env-required:"false"`
}

var Cfg Config

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Cfg.API_KEY = os.Getenv("API_KEY")
	Cfg.TRANSCRIPTION_ENDPOINT = os.Getenv("TRANSCRIPTION_ENDPOINT")
	Cfg.MODEL = os.Getenv("MODEL")
	Cfg.SERVER_URL = os.Getenv("SERVER_URL")
	Cfg.RESPONSE_FORMAT = os.Getenv("RESPONSE_FORMAT")
	Cfg.TIMESTAMP_GRANULARITIES = os.Getenv("TIMESTAMP_GRANULARITIES")
	Cfg.COMPLETION_ENDPOINT = os.Getenv("COMPLETION_ENDPOINT")
	Cfg.COMPLETION_MESSAGE = os.Getenv("COMPLETION_MESSAGE")
}

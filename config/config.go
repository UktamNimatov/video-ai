package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	API_KEY                string `env:"API_KEY" env-required:"false"`
	TRANSCRIPTION_ENDPOINT string `env:"TRANSCRIPTION_ENDPOINT" env-required:"false"`
	MODEL                  string `env:"MODEL" env-required:"false"`
	SERVER_URL             string `env:"SERVER_URL" env-required:"false"`
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
}

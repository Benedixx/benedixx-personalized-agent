package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Config *Configuration

type Configuration struct {
	OpenRouterKey string
	PrimaryLLM    string
	ReasoningLLM  string
	SmallLLM      string
	OllamaURL     string
}

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		Warn("error loading .env file")
		return err
	}

	// init env secrets
	Config = &Configuration{
		OpenRouterKey: os.Getenv("OPENROUTER_API_KEY"),
		PrimaryLLM:    os.Getenv("PRIMARY_LLM_MODEL"),
		ReasoningLLM:  os.Getenv("REASONING_LLM_MODEL"),
		SmallLLM:      os.Getenv("SMALL_LLM_MODEL"),
		OllamaURL:     os.Getenv("OLLAMA_URL"),
	}

	return nil
}

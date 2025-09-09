package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Config *Configuration

type Configuration struct {
	OpenRouterKey  string
	PrimaryLLM     string
	ReasoningLLM   string
	SmallLLM       string
	OllamaURL      string
	EmbeddingModel string
	Database       *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		Warn("error loading .env file")
		return err
	}

	// init env secrets
	Config = &Configuration{
		OpenRouterKey:  os.Getenv("OPENROUTER_API_KEY"),
		PrimaryLLM:     os.Getenv("PRIMARY_LLM_MODEL"),
		ReasoningLLM:   os.Getenv("REASONING_LLM_MODEL"),
		SmallLLM:       os.Getenv("SMALL_LLM_MODEL"),
		OllamaURL:      os.Getenv("OLLAMA_URL"),
		EmbeddingModel: os.Getenv("EMBEDDING_MODEL"),
		Database: &DatabaseConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASS"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
	}

	return nil
}

package core

import (
	"benedixx-personalized-agent/src/config"
	"benedixx-personalized-agent/src/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	httpClient = &http.Client{Timeout: 60 * time.Second}
)

func ChatCompletion(model string,
	messages []map[string]interface{},
	stream bool,
	options map[string]interface{}) (interface{}, error) {

	reqBody := dto.ChatRequest{
		Model:    model,
		Messages: messages,
		Stream:   stream,
		Options:  options,
	}

	jsonData, err := json.Marshal(reqBody)

	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", config.Config.OllamaURL+"/api/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	config.Log.Info("Sending request to Ollama API")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	var respBody dto.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	return respBody.Message["content"], nil
}

package core

import (
	"benedixx-personalized-agent/src/config"
	"benedixx-personalized-agent/src/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GenerateEmbedding(text []string) ([][]float64, error) {
	reqBody := dto.EmbedRequest{
		Model:  config.Config.EmbeddingModel,
		Inputs: text,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Call the embedding API
	resp, err := http.Post(config.Config.OllamaURL+"/api/embed", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to generate embedding: %s", resp.Status)
	}

	var embeddingResponse dto.EmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&embeddingResponse); err != nil {
		return nil, err
	}

	if len(embeddingResponse.Embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned from API - model: %s, input count: %d", embeddingResponse.Model, len(text))
	}

	return embeddingResponse.Embeddings, nil
}

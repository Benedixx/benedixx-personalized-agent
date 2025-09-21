package service

import (
	"benedixx-personalized-agent/src/config"
	"benedixx-personalized-agent/src/core"
	"benedixx-personalized-agent/src/dto"
	"encoding/json"
	"fmt"
	"strings"
)

func ExtractEntities(text string) ([]dto.EntityData, error) {
	systemPrompt := `Extract named entities from the provided text. ONLY respond with list of entities`

	userPrompt := fmt.Sprintf("Extract entities from this text:\n\n%s", text)

	messages := []map[string]interface{}{
		{"role": "system", "content": systemPrompt},
		{"role": "user", "content": userPrompt},
	}

	options := map[string]interface{}{
		"max_tokens":  800,
		"temperature": 0.1,
	}

	response, err := core.ChatCompletion(config.Config.SmallLLM, messages, false, options)
	if err != nil {
		return nil, err
	}

	responseStr := response.(string)

	// Clean up response
	responseStr = strings.TrimSpace(responseStr)
	responseStr = strings.TrimPrefix(responseStr, "```json")
	responseStr = strings.TrimPrefix(responseStr, "```")
	responseStr = strings.TrimSuffix(responseStr, "```")
	responseStr = strings.TrimSpace(responseStr)

	// Validate that response starts and ends with square brackets
	if !strings.HasPrefix(responseStr, "[") || !strings.HasSuffix(responseStr, "]") {
		return nil, fmt.Errorf("invalid JSON format: response should be a JSON array")
	}

	var entities []dto.EntityData
	err = json.Unmarshal([]byte(responseStr), &entities)
	if err != nil {
		config.Log.Error("Failed to parse entity extraction response", "response", responseStr, "error", err)
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Filter out any entities with placeholder names
	var validEntities []dto.EntityData
	for _, entity := range entities {
		// Skip entities with placeholder or example names
		if entity.Name != "entity_name" &&
			entity.Name != "actual_entity_name" &&
			strings.TrimSpace(entity.Name) != "" {
			validEntities = append(validEntities, entity)
		}
	}

	return validEntities, nil
}

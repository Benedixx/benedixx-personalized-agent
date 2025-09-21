package service

import (
	"benedixx-personalized-agent/src/core"
	"benedixx-personalized-agent/src/database"
	"benedixx-personalized-agent/src/dto"
)

func RetrieveRelevantChunks(query string, topK int) ([]dto.ChunkResult, error) {
	embedding, err := core.GenerateEmbedding([]string{query})
	if err != nil {
		return nil, err
	}
	chunks, err := database.GetRelevantChunks(embedding[0], topK)
	if err != nil {
		return nil, err
	}
	return chunks, nil
}

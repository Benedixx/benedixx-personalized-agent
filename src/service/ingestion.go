package service

import (
	"benedixx-personalized-agent/src/config"
	"benedixx-personalized-agent/src/core"
	"benedixx-personalized-agent/src/database"
	"benedixx-personalized-agent/src/dto"
	"fmt"
	"os"

	"github.com/oliverpool/unipdf/v3/extractor"
	"github.com/oliverpool/unipdf/v3/model"
)

func ReadPDF(path string, fileMetadata dto.FileMetadata) (map[string]interface{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader, err := model.NewPdfReader(f)
	if err != nil {
		return nil, err
	}

	np, _ := reader.GetNumPages()
	documentData := make(map[string]interface{})
	totalExtractedText := ""

	for i := 1; i <= np; i++ {
		page, _ := reader.GetPage(i)
		ex, _ := extractor.New(page)
		txt, err := ex.ExtractText()
		if err != nil {
			config.Log.Error("failed to extract text", "error", err, "page", i)
			continue
		}
		if len(txt) == 0 {
			config.Log.Warn("empty page", "page", i)
			continue
		}

		totalExtractedText += txt

		documentData[fmt.Sprintf("page_%d", i)] = map[string]interface{}{
			"text": txt,
			"page": i,
		}
	}

	if len(totalExtractedText) == 0 {
		return nil, fmt.Errorf("failed to extract any text from document")
	}

	result := map[string]interface{}{
		"metadata": fileMetadata,
		"pages":    documentData,
	}

	return result, nil
}

func ChunkText(text string, chunkSize int, overlap int) []string {
	var chunks []string
	dataLen := len(text)

	for start := 0; start < dataLen; {
		end := start + chunkSize
		if end > dataLen {
			end = dataLen
		}
		if end < dataLen {
			for end > start && text[end-1] != '.' && text[end-1] != '!' && text[end-1] != '?' {
				end--
			}
		}
		chunks = append(chunks, text[start:end])

		if end == dataLen {
			break
		}
		start = end - overlap
		if start < 0 {
			start = 0
		}
	}

	return chunks
}

func IngestDocument(path string, fileMetadata dto.FileMetadata) error {
	docData, err := ReadPDF(path, fileMetadata)
	if err != nil {
		return err
	}

	pages := docData["pages"].(map[string]interface{})

	var allChunks []dto.ChunkData

	for _, v := range pages {
		page := v.(map[string]interface{})
		txt := page["text"].(string)

		chunks := ChunkText(txt, config.ChunkSize, config.ChunkOverlap)

		// batch embeddings
		embeddingResults, err := core.GenerateEmbedding(chunks)
		if err != nil {
			config.Log.Error("failed to generate embeddings ", err)
			return err
		}

		for j, chunk := range chunks {
			embedding := embeddingResults[j]
			chunkData := dto.ChunkData{
				Text:       chunk,
				Embedding:  embedding,
				PageNumber: page["page"].(int),
			}
			allChunks = append(allChunks, chunkData)
		}
	}

	err = database.UpsertDocumentWithChunks(fileMetadata, allChunks)
	if err != nil {
		config.Log.Error("failed to upsert document with chunks", "error", err)
		return err
	}

	return nil
}

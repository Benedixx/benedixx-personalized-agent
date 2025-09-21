package database

import (
	"benedixx-personalized-agent/src/dto"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// float64SliceToVector converts a []float64 to a PostgreSQL vector string format
func float64SliceToVector(embedding []float64) string {
	if len(embedding) == 0 {
		return "[]"
	}

	var parts []string
	for _, v := range embedding {
		parts = append(parts, fmt.Sprintf("%.6f", v))
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func UpsertDocument(tx *sqlx.Tx, file_metadata dto.FileMetadata) (int, error) {
	var documentId int
	err := tx.QueryRow("INSERT INTO documents (document_name, document_author, document_year) VALUES ($1, $2, $3) RETURNING id",
		file_metadata.Title,
		file_metadata.Author,
		file_metadata.Year).Scan(&documentId)

	if err != nil {
		return 0, fmt.Errorf("failed to insert document: %w", err)
	}

	return documentId, nil
}

func UpsertChunk(tx *sqlx.Tx, documentId int, chunkText string, embedding []float64, pageNumber int) error {
	vectorStr := float64SliceToVector(embedding)
	_, err := tx.Exec("INSERT INTO chunks (document_id, chunk_text, embedding, page_number) VALUES ($1, $2, $3, $4)",
		documentId,
		chunkText,
		vectorStr,
		pageNumber)
	if err != nil {
		return fmt.Errorf("failed to insert chunk: %w", err)
	}

	return nil
}

func UpsertDocumentWithChunks(fileMetadata dto.FileMetadata, chunks []dto.ChunkData) error {
	db := GetDB()
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert document
	documentId, err := UpsertDocument(tx, fileMetadata)
	if err != nil {
		return err
	}

	// Insert all chunks
	for _, chunk := range chunks {
		err = UpsertChunk(tx, documentId, chunk.Text, chunk.Embedding, chunk.PageNumber)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func GetRelevantChunks(embedding []float64, topK int) ([]dto.ChunkResult, error) {
	db := GetDB()
	vectorStr := float64SliceToVector(embedding)
	query := `
        SELECT chunk_text, document_id, page_number
        FROM chunks
        ORDER BY embedding <-> $1::vector
        LIMIT $2;
    `
	var chunks []dto.ChunkResult
	err := db.Select(&chunks, query, vectorStr, topK)
	if err != nil {
		return nil, fmt.Errorf("failed to get relevant chunks: %w", err)
	}
	return chunks, nil
}

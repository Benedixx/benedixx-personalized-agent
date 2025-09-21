package dto

import "mime/multipart"

// request dto for ingestion
type UploadPDFRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type FileMetadata struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"required"`
}

type ChunkData struct {
	Text       string    `json:"text"`
	Embedding  []float64 `json:"embedding"`
	PageNumber int       `json:"page_number"`
}

type IngestDocRequest struct {
	File     *multipart.FileHeader `form:"file" binding:"required"`
	Metadata FileMetadata          `json:"metadata"`
}

// request dto for model inference
type ChatRequest struct {
	Model    string                   `json:"model" binding:"required"`
	Messages []map[string]interface{} `json:"messages" binding:"required"`
	Stream   bool                     `json:"stream" default:"false"`
	Options  map[string]interface{}   `json:"options"  omitempty:"true"`
}

// request dto for model embedding inference
type EmbedRequest struct {
	Model  string   `json:"model" binding:"required"`
	Inputs []string `json:"input" binding:"required"`
}

type EntityData struct {
	Name              string `json:"name"`
	Type              string `json:"type"`
	EntityDescription string `json:"description"`
}

type RelationData struct {
	Source     string `json:"source"`
	Target     string `json:"target"`
	Relation   string `json:"relation"`
	ChunkIndex int    `json:"chunk_index"`
}

type ChunkResult struct {
	ChunkText  string `db:"chunk_text"`
	DocumentID int    `db:"document_id"`
	PageNumber int    `db:"page_number"`
}

type RetrieveRequest struct {
	Query string `form:"query" binding:"required"`
	TopK  *int   `form:"top_k"`
}

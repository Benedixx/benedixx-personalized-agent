package dto

import "mime/multipart"

// request dto for ingestion
type UploadPDFRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type FileMetadata struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type IngestDocRequest struct {
	File     *multipart.FileHeader `form:"file" binding:"required"`
	Metadata string                `form:"metadata" binding:"required"`
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

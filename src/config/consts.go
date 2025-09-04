package config

const (
	// file ingestion consts
	MaxDocsFileSize = 50 * 1024 * 1024 // 50mb
	MaxDocsPages    = 1000
	AllowedFileExts = ".pdf"

	// API Limits
	DefaultTimeout = 30
	MaxRetries     = 3

	// chunking consts
	ChunkSize    = 5000 // chars
	ChunkOverlap = 500  // chars
)

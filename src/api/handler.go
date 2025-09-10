package api

import (
	"benedixx-personalized-agent/src/core"
	"benedixx-personalized-agent/src/dto"
	"benedixx-personalized-agent/src/service"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ChatHandler(c *gin.Context) {
	var request dto.ChatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := core.ChatCompletion(request.Model, request.Messages, request.Stream, request.Options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func EmbedHandler(c *gin.Context) {
	var request dto.EmbedRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := core.GenerateEmbedding(request.Inputs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func IngestDoc(c *gin.Context) {
	// file upload
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// file meatdata
	metadataStr := c.PostForm("metadata")
	if metadataStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "metadata is required"})
		return
	}

	var metadata dto.FileMetadata
	if err := json.Unmarshal([]byte(metadataStr), &metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid metadata format"})
		return
	}

	//validate metadata
	validate := validator.New()
	if err := validate.Struct(metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempPath := filepath.Join("temp", file.Filename)
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}
	defer os.Remove(tempPath)

	err = service.IngestDocument(tempPath, metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

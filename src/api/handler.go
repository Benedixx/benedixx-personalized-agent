package api

import (
	"benedixx-personalized-agent/src/core"
	"benedixx-personalized-agent/src/dto"
	"net/http"

	"github.com/gin-gonic/gin"
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

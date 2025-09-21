package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/chat", ChatHandler)
	r.POST("/embed", EmbedHandler)
	r.POST("/ingest", IngestDoc)
	r.POST("/retrieve", RetrieveHandler)

	return r
}

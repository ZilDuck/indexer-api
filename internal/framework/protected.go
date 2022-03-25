package framework

import (
	"github.com/ZilDuck/indexer-api/internal/auth"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Protected(c *gin.Context) {
	apiKey := helpers.ApiKey(c)
	if apiKey == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing API key", "status": http.StatusUnauthorized})
		return
	}

	_, err := auth.GetApiClient(apiKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key", "status": http.StatusUnauthorized})
		return
	}


	zap.L().With(zap.String("apiKey", apiKey)).Info("Authenticated")

	c.Next()
}

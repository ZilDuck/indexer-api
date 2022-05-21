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

	client, err := auth.GetClientByApiKey(apiKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key", "status": http.StatusUnauthorized})
		return
	}

	zap.L().With(zap.String("client", client.Username)).Info("Authenticated")

	c.Next()
}

func ProtectedAdmin(c *gin.Context) {
	apiKey := helpers.ApiKey(c)
	if apiKey == "" {
		zap.L().Error("ApiKey not found")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
		return
	}

	client, err := auth.GetClientByApiKey(apiKey)
	if err != nil {
		zap.L().Error("ApiKey not valid")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
		return
	}

	if !client.IsAdmin() {
		zap.L().Error("ApiKey not admin")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
		return
	}

	//for _, adminId := range config.Get().AdminIds {
	//	if client.ID.String() == adminId {
	//		zap.L().With(zap.String("client", client.Username)).Info("Authenticated as Admin")
	//		c.Next()
	//		return
	//	}
	//}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
	return
}

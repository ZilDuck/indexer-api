package resource

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func handleError(c *gin.Context, err error, msg string, status int) {
	if err != nil {
		zap.L().With(zap.Error(err)).Error(msg)
	}
	c.AbortWithStatusJSON(status, gin.H{"message": msg, "status": status})
}

func jsonResponse(c *gin.Context, object interface{}) {
	if c.GetHeader("Cache-Control") == "" {
		c.Header("Cache-Control", "max-age=60")
	}

	c.JSON(200, object)
}
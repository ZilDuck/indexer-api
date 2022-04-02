package framework

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SetReleaseMode(debug bool) {
	if debug {
		zap.S().Infof("Release Mode: %s", gin.DebugMode)
		gin.SetMode(gin.DebugMode)
	} else {
		zap.S().Infof("Release Mode: %s", gin.ReleaseMode)
		gin.SetMode(gin.ReleaseMode)
	}
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

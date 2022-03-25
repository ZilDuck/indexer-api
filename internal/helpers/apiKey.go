package helpers

import (
	"github.com/gin-gonic/gin"
)

func ApiKey(c *gin.Context) string {
	if c.Request.Header.Get("X-API-KEY") != "" {
		return c.Request.Header.Get("X-API-KEY")
	}

	return c.DefaultQuery("apikey", "")
}

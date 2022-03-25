package helpers

import "github.com/gin-gonic/gin"

func ApiKey(c *gin.Context) string {
	return c.Request.Header.Get("X-API-KEY")
}

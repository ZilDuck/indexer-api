package resource

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	MAINNET = "mainnet"
	TESTNET = "testnet"
)

func network(c *gin.Context) (n string) {
	networkQuery, _ := c.GetQuery("network")
	if networkQuery == "https://dev-api.zilliqa.com" || networkQuery == "testnet" {
		return TESTNET
	}

	return MAINNET
}

func handleError(c *gin.Context, err error, msg string, status int) {
	zap.L().With(zap.Error(err)).Error(msg)
	c.AbortWithStatusJSON(status, gin.H{"message": msg, "status": status})
}

func jsonResponse(c *gin.Context, object interface{}) {
	c.Header("Cache-Control", "no-cache")
	c.JSON(200, object)
}
package resource

import (
	"github.com/Zilliqa/gozilliqa-sdk/bech32"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
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

func getAddress(address string) string {
	address = strings.ToLower(address)

	if address[0:3] == "zil" {
		if fromBech32, err := bech32.FromBech32Addr(address); err == nil {
			address = "0x"+strings.ToLower(fromBech32)
		}
	}
	zap.L().Info("Using address "+address)
	return address
}
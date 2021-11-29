package resource

import (
	"github.com/gin-gonic/gin"
)

const (
	MAINNET = "mainnet"
	TESTNET = "testnet"
)

func network(c *gin.Context) (n string) {
	networkQuery, _ := c.GetQuery("network")
	if networkQuery == "https://dev-api.zilliqa.com" {
		return TESTNET
	}

	return MAINNET
}

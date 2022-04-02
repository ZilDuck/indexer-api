package helpers

import "github.com/gin-gonic/gin"

const (
	Mainnet = "mainnet"
	Testnet = "testnet"
)

func Network(c *gin.Context) (n string) {
	networkQuery, _ := c.GetQuery("network")
	if networkQuery == "https://dev-api.zilliqa.com" || networkQuery == "testnet" {
		return Testnet
	}

	return Mainnet
}

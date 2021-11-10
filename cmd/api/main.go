package main

import (
	"fmt"
	"github.com/dantudor/zilkroad-txapi/generated/dic"
	"github.com/dantudor/zilkroad-txapi/internal/config"
	"github.com/dantudor/zilkroad-txapi/internal/framework"
	"github.com/dantudor/zilkroad-txapi/internal/resource"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/dingo/v3"
	"net/http"
)

var container *dic.Container

func main() {
	config.Init()
	container, _ = dic.NewContainer(dingo.App)

	framework.SetReleaseMode(config.Get().Debug)

	r := gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(framework.Cors())
	r.Use(framework.Options)
	r.Use(framework.ErrorHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to ZilkRoad NFT-API!")
	})

	nftResource := resource.NewNftResource(container.GetNftService())
	r.GET("/nfts/:ownerAddr", nftResource.GetNftsOwnedByAddress)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Resource not found"})
	})

	_ = r.Run(fmt.Sprintf(":%d", config.Get().Port))
}

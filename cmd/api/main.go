package main

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/generated/dic"
	"github.com/ZilDuck/indexer-api/internal/audit"
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/ZilDuck/indexer-api/internal/resource"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var container *dic.Container

func main() {
	config.Init()
	audit.Init(config.Get().AuditDir)

	container, _ = dic.NewContainer()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), setupRouter()); err != nil {
		zap.L().With(zap.Error(err)).Fatal("Failed to start API")
	}
}

func setupRouter() *gin.Engine {
	framework.SetReleaseMode(config.Get().Debug)

	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(framework.Cors)
	r.Use(framework.Options)
	r.Use(framework.ErrorHandler)

	apigroup := r.Group("/", framework.RateLimiter)

	contractResource := resource.NewContractResource(container.GetContractRepository(), container.GetNftRepository())
	apigroup.GET("/contract/:contractAddr", contractResource.GetContract)
	apigroup.GET("/contract/:contractAddr/code", contractResource.GetCode)
	apigroup.GET("/contract/:contractAddr/attributes", contractResource.GetAttributes)

	nftResource := resource.NewNftResource(container.GetNftRepository(), container.GetActionRepository(), container.GetMessenger(), container.GetMetadataService())
	apigroup.GET("/nft/:contractAddr", nftResource.GetContractNfts)
	apigroup.GET("/nft/:contractAddr/:tokenId", nftResource.GetContractNft)
	apigroup.GET("/nft/:contractAddr/:tokenId/refresh", nftResource.RefreshMetadata)
	apigroup.GET("/nft/:contractAddr/:tokenId/metadata", nftResource.GetContractNftMetadata)
	apigroup.GET("/nft/:contractAddr/:tokenId/actions", nftResource.GetContractNftActions)

	apigroup.GET("/address/:ownerAddr/nft", nftResource.GetNftsOwnedByAddress)
	apigroup.GET("/address/:ownerAddr/contract", contractResource.GetContractsOwnedByAddress)

	apigroup.GET("/health", resource.NewHealthResource(container.GetElastic()).HealthCheck)

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("Welcome to the NFT index API"))
	})
	r.GET("/loaderio-b8e545b85c125048324e20015fd1fc45.txt", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("loaderio-b8e545b85c125048324e20015fd1fc45"))
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
	})

	return r
}
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

	auditResource := resource.NewAuditResource(container.GetAuditRepository())
	r.GET("/audit/status", auditResource.GetStatus)
	r.GET("/audit/log/:month", auditResource.GetLogsForDate)

	contractResource := resource.NewContractResource(container.GetContractRepository(), container.GetContractStateRepository(), container.GetNftRepository())
	r.GET("/contract/:contractAddr", audit.Handler, contractResource.GetContract)
	r.GET("/contract/:contractAddr/code", audit.Handler, contractResource.GetCode)
	r.GET("/contract/:contractAddr/attributes", audit.Handler, contractResource.GetAttributes)
	r.GET("/contract/:contractAddr/state", audit.Handler, contractResource.GetState)

	nftResource := resource.NewNftResource(container.GetNftRepository(), container.GetActionRepository(), container.GetMessenger(), container.GetMetadataService())
	r.GET("/nft/:contractAddr", audit.Handler, nftResource.GetContractNfts)
	r.GET("/nft/:contractAddr/:tokenId", audit.Handler, nftResource.GetContractNft)
	r.GET("/nft/:contractAddr/:tokenId/refresh", audit.Handler, nftResource.RefreshMetadata)
	r.GET("/nft/:contractAddr/:tokenId/metadata", audit.Handler, nftResource.GetContractNftMetadata)
	r.GET("/nft/:contractAddr/:tokenId/actions", audit.Handler, nftResource.GetContractNftActions)

	r.GET("/address/:ownerAddr/nft", audit.Handler, nftResource.GetNftsOwnedByAddress)
	r.GET("/address/:ownerAddr/contract", audit.Handler, contractResource.GetContractsOwnedByAddress)

	r.GET("/health", resource.NewHealthResource(container.GetElastic()).HealthCheck)

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("Welcome to the NFT index API"))
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
	})

	return r
}
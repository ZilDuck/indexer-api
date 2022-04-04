package main

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/generated/dic"
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

	container, _ = dic.NewContainer()
	container.GetAuthService().LoadClients()

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

	protect := r.Group("/", framework.Protected)

	auditResource := resource.NewAuditResource(container.GetAuditRepository())
	protect.GET("/audit/status/:month", auditResource.GetStatus)
	protect.GET("/audit/log/:month", auditResource.GetLogsForDate)

	contractResource := resource.NewContractResource(container.GetContractRepository(), container.GetContractStateRepository(), container.GetNftRepository())
	protect.GET("/contract/:contractAddr", contractResource.GetContract)
	protect.GET("/contract/:contractAddr/code", contractResource.GetCode)
	protect.GET("/contract/:contractAddr/attributes", contractResource.GetAttributes)
	protect.GET("/contract/:contractAddr/state", contractResource.GetState)

	nftResource := resource.NewNftResource(container.GetNftRepository(), container.GetActionRepository(), container.GetMessenger(), container.GetMetadataService())
	protect.GET("/nft/:contractAddr", nftResource.GetContractNfts)
	protect.GET("/nft/:contractAddr/:tokenId", nftResource.GetContractNft)
	protect.GET("/nft/:contractAddr/:tokenId/refresh", nftResource.RefreshMetadata)
	protect.GET("/nft/:contractAddr/:tokenId/metadata", nftResource.GetContractNftMetadata)
	protect.GET("/nft/:contractAddr/:tokenId/actions", nftResource.GetContractNftActions)

	protect.GET("/address/:ownerAddr/nft", nftResource.GetNftsOwnedByAddress)
	protect.GET("/address/:ownerAddr/contract", contractResource.GetContractsOwnedByAddress)

	protect.GET("/health", resource.NewHealthResource(container.GetElastic()).HealthCheck)

	//adminResource := resource.NewAdminResource(container.GetAuthService())
	//adminRoute := r.Group("/admin", framework.ProtectedAdmin)
	//adminRoute.GET("/client", adminResource.GetClients)
	//adminRoute.POST("/client", adminResource.CreateClient)

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("Welcome to the NFT index API"))
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Resource not found"})
	})

	return r
}
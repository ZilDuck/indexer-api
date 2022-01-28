package main

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/generated/dic"
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/ZilDuck/indexer-api/internal/resource"
	"github.com/apex/gateway"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var container *dic.Container

func main() {
	config.Init()
	container, _ = dic.NewContainer()

	if inLambda() {
		zap.L().Info("Running lambda")
		if err := gateway.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), setupRouter()); err != nil {
			zap.L().With(zap.Error(err)).Fatal("Failed to execute gateway")
		}
	} else {
		zap.L().Info("Running naked")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Get().Port), setupRouter()); err != nil {
			zap.L().With(zap.Error(err)).Fatal("Failed to execute lambda")
		}
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

	contractResource := resource.NewContractResource(container.GetContractRepository())
	r.GET("/contract/:contractAddr", contractResource.GetContract)

	nftResource := resource.NewNftResource(container.GetNftRepository(), container.GetMetadataService())
	r.GET("/nft/:contractAddr", nftResource.GetContractNfts)
	r.GET("/nft/:contractAddr/:tokenId", nftResource.GetContractNft)
	r.GET("/nft/:contractAddr/:tokenId/metadata", nftResource.GetContractNftMetadata)
	r.GET("/nft/:contractAddr/:tokenId/asset", nftResource.GetContractNftAsset)

	r.GET("/wallets/:ownerAddr", nftResource.GetNftsOwnedByAddress)

	return r
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

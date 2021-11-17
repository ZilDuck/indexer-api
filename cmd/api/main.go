package main

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/generated/dic"
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/ZilDuck/indexer-api/internal/resource"
	"github.com/apex/gateway"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	middleware "github.com/s12i/gin-throttle"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var container *dic.Container

func main() {
	config.Init()
	container, _ = dic.NewContainer()

	if config.Get().Cache && config.Get().Subscribe {
		go container.GetCacheValidator().Subscribe()
	}

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
	r.Use(framework.Cors())
	r.Use(framework.Options)
	r.Use(framework.ErrorHandler)

	r.Use(middleware.Throttle(config.Get().Throttle.MaxEventsPerSec, config.Get().Throttle.MaxBurstSize))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to ZilkRoad NFT-API!")
	})

	nftResource := resource.NewNftResource(container.GetNftService(), container.GetCacheStore())

	if config.Get().Cache {
		r.GET("/cache-clear", func(c *gin.Context) {
			if err := container.GetCacheStore().Flush(); err != nil {
				zap.L().With(zap.Error(err)).Error("Failed to flush cache")
			}
			c.String(http.StatusOK, "")
		})

		r.GET("/nfts/:ownerAddr",
			cache.CachePageWithoutHeader(container.GetCacheStore(), config.Get().CacheDefaultExpiration, nftResource.GetNftsOwnedByAddress),
		)
	} else {
		r.GET("/nfts/:ownerAddr", nftResource.GetNftsOwnedByAddress)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "Resource not found"})
	})

	return r
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

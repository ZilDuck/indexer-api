package main

import (
	"context"
	"fmt"
	"github.com/ZilDuck/indexer-api/generated/dic"
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/apex/gateway"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
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
		if err := gateway.ListenAndServe(fmt.Sprintf(":%d", config.Get().DashboardPort), setupRouter()); err != nil {
			zap.L().With(zap.Error(err)).Fatal("Failed to execute gateway")
		}
	} else {
		zap.L().Info("Running naked")
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Get().DashboardPort), setupRouter()); err != nil {
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

	r.NoRoute(func(c *gin.Context) {
		container.GetElastic().Client.Index()

		resp, err := container.GetElastic().Client.PerformRequest(context.Background(), elastic.PerformRequestOptions{
			Method:      c.Request.Method,
			Path:        c.Request.RequestURI,
			Params:      nil,
			Body:        c.Request.Body,
			ContentType: "application/x-ndjson",
			Retrier:     nil,
			Headers:     nil,
		})

		if err != nil {
			c.JSON(500, gin.H{"code": 500, "message": err.Error()})
			return
		}

		c.JSON(resp.StatusCode, resp.Body)
	})

	return r
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

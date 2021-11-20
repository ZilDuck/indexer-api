package main

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/generated/dic"
	"github.com/ZilDuck/indexer-api/internal/config"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/apex/gateway"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sha1sum/aws_signing_client"
	"go.uber.org/zap"
	"io/ioutil"
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
		awsConfig := config.Get().Aws
		creds := credentials.NewStaticCredentials(awsConfig.AccessKey, awsConfig.SecretKey, awsConfig.Token)
		awsClient, err := aws_signing_client.New(v4.NewSigner(creds), nil, "es", awsConfig.Region)
		if err != nil {
			panic(err)
		}

		baseUrl := "http://search-zilkroad-index-wlmccxpkwz6ps7sohqqjvdphqa.us-east-1.es.amazonaws.com"
		req, err := http.NewRequest(c.Request.Method, fmt.Sprintf("%s%s", baseUrl, c.Request.RequestURI), c.Request.Body)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		resp, err := awsClient.Do(req)
		if err != nil {
			zap.L().Error(err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			zap.L().Error(err.Error())
			c.JSON(500, gin.H{"code": 500, "message": err.Error()})
			return
		}

		for key, value := range resp.Header {
			for _, v := range value {
				c.Header(key, v)
			}
		}
		zap.L().Info(string(body))

		c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
	})

	return r
}

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

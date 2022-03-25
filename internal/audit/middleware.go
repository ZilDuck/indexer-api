package audit

import (
	"encoding/json"
	"github.com/ZilDuck/indexer-api/internal/entity"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"sync"
)

var mu sync.Mutex

func Handler(c *gin.Context) {
	apiKey := helpers.ApiKey(c)
	uri := c.Request.RequestURI

	zap.L().With(zap.String("apiKey", apiKey), zap.String("request", uri)).Info("Audit")

	audit := entity.Audit{
		ApiKey:     apiKey,
		Request:    uri,
		Network:    helpers.Network(c),
	}

	if c.Request.Header.Get("Cf-Connecting-Ip") != "" {
		audit.RemoteAddr = c.Request.Header.Get("Cf-Connecting-Ip")
	} else if c.Request.Header.Get("X-Forwarded-For") != "" {
		audit.RemoteAddr = c.Request.Header.Get("X-Forwarded-For")
	} else {
		audit.RemoteAddr = c.Request.RemoteAddr
	}

	data, err := json.Marshal(audit)
	if err != nil {
		zap.L().With(zap.String("apiKey", apiKey), zap.String("request", uri)).Error("Failed to generate audit")
	}
	go write(data)

	c.Next()
}

func write(data []byte) {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.OpenFile(getAuditFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		zap.L().With(zap.Error(err)).Error("Failed to prepare audit")
	}
	defer f.Close()

	if _, err := f.WriteString(string(data)+"\n"); err != nil {
		zap.L().With(zap.Error(err)).Error("Failed to write audit")
	}
}

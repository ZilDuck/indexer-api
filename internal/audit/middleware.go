package audit

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"sync"
)

var mu sync.Mutex

func Audit(c *gin.Context) {
	apiKey := c.Request.Header.Get("X-API-KEY")
	uri := c.Request.RequestURI

	zap.L().With(zap.String("apiKey", apiKey), zap.String("request", uri)).Info("Audit")

	audit := map[string]interface{}{}
	audit["apiKey"] = apiKey
	audit["request"] = uri
	audit["remoteAddr"] = c.Request.RemoteAddr
	audit["referer"] = c.Request.Referer()

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

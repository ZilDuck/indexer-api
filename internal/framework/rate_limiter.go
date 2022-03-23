package framework

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/audit"
	"github.com/ZilDuck/indexer-api/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

var limiters = make(map[string]*ratelimit.Bucket)

func RateLimiter(c *gin.Context) {
	apiKey := c.Request.Header.Get("X-API-KEY")

	client, err := auth.GetApiClient(apiKey)
	if err != nil {
		if errors.Is(err, auth.ErrNoClientFound) {
			c.Next()
			return
		}

		if apiKey == "" {
			c.Header("Content-Type", "application/json")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing API Key", "status": http.StatusBadRequest})
			return
		}

		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid API Key", "status": http.StatusUnauthorized})
		return
	}

	if _, exists := limiters[apiKey]; !exists {
		limiters[apiKey] = ratelimit.NewBucketWithQuantum(time.Second * client.Duration, client.Capacity, client.Capacity)
	}

	limiter := limiters[apiKey]
	if limiter.TakeAvailable(1) == 0 {
		c.Header("Content-Type", "application/json")
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Too many requests", "status": http.StatusTooManyRequests})
		return
	}

	audit.Audit(c)

	c.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
	c.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
	c.Next()
}
package resource

import (
	"context"
	"github.com/ZilDuck/indexer-api/internal/dto"
	"github.com/ZilDuck/indexer-api/internal/elastic_search"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthResource struct {
	elastic elastic_search.Index
}

func NewHealthResource(elastic elastic_search.Index) HealthResource {
	return HealthResource{elastic}
}

func (r HealthResource) HealthCheck(c *gin.Context) {
	health := dto.HealthCheck{}
	health.Up(r.elasticSearchHealth())

	c.Header("Cache-Control", "no-store")

	if health.Healthy() {
		c.JSON(http.StatusOK, health)
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, health)
	}
}

func (r HealthResource) elasticSearchHealth() bool {
	status := false
	elasticHealth, err := r.elastic.Client.CatHealth().Do(context.Background())
	if err == nil {
		for _, node := range elasticHealth {
			if node.Status == "green" {
				status = true
				break
			}
		}
	}

	return status
}

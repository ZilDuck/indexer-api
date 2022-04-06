package resource

import (
	"encoding/json"
	"github.com/ZilDuck/indexer-api/internal/request"
	"github.com/gin-gonic/gin"
	"math"
)

type paginationHeader struct {
	Size          int `json:"size"`
	Page          int `json:"page"`
	TotalPages    int `json:"total_pages"`
	TotalElements int `json:"total_elements"`
}

func paginator(c *gin.Context, total int64, pagination *request.Pagination) {
	pages := int(math.Ceil(float64(total) / float64(pagination.Size)))
	if pages == 0 || total == 0 {
		pages = 1
	}

	header, _ := json.Marshal(paginationHeader{
		Size:          pagination.Size,
		Page:          pagination.Page,
		TotalPages:    pages,
		TotalElements: int(total),
	})

	c.Writer.Header().Set("X-Pagination", string(header))
}
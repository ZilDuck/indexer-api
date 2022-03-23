package resource

import (
	"encoding/json"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/gin-gonic/gin"
	"math"
)

type paginationHeader struct {
	Limit         uint64 `json:"limit"`
	Offset        uint64 `json:"offset"`
	TotalPages    int    `json:"total_pages"`
	TotalElements int    `json:"total_elements"`
}

func paginator(c *gin.Context, total int64, pagination framework.Pagination) {
	pages := int(math.Ceil(float64(total) / float64(pagination.Size)))
	if pages == 0 || total == 0 {
		pages = 1
	}

	header, _ := json.Marshal(paginationHeader{
		Limit:         pagination.Size,
		Offset:        pagination.Offset,
		TotalPages:    pages,
		TotalElements: int(total),
	})

	c.Writer.Header().Set("X-Pagination", string(header))
}
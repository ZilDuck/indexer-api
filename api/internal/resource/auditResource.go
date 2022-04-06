package resource

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/ZilDuck/indexer-api/internal/mapper"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/ZilDuck/indexer-api/internal/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AuditResource struct {
	auditRepo repository.AuditRepository
}

func NewAuditResource(auditRepo repository.AuditRepository) AuditResource {
	return AuditResource{auditRepo}
}

func (r AuditResource) GetStatus(c *gin.Context) {
	month, err := getMonthOffset(c)
	if err != nil {
		handleError(c, err, err.Error(), http.StatusBadRequest)
	}

	total, err := r.auditRepo.CountByDateAndApiKey(*month, helpers.ApiKey(c))
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get audit count"), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"requestCount": total})
	c.Header("Cache-Control", "no-store")
}

func (r AuditResource) GetLogsForDate(c *gin.Context) {
	req := request.NewPaginatedRequest(c)

	month, err := getMonthOffset(c)
	if err != nil {
		handleError(c, err, err.Error(), http.StatusBadRequest)
	}

	audits, total, err := r.auditRepo.GetByDateAndApiKey(*month, helpers.ApiKey(c), req.Pagination.Size, req.Pagination.Offset)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get audit logs"), http.StatusInternalServerError)
		return
	}

	paginator(c, total, req.Pagination)
	jsonResponse(c, mapper.AuditsToDtos(audits))
	c.Header("Cache-Control", "no-store")
}

func getMonthOffset(c *gin.Context) (*time.Time, error) {
	i, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		return nil, errors.New("month must be a number between 1 and -11")
	}
	if i > 0 {
		return nil, errors.New("month cannot be greater than 0")
	}
	if i < -11 {
		return nil, errors.New("month cannot be less than -11")
	}

	month := time.Now().AddDate(0, i, 0)

	return &month, nil
}
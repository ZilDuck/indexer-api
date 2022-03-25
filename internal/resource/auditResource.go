package resource

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/ZilDuck/indexer-api/internal/repository"
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

}

func (r AuditResource) GetLogsForDate(c *gin.Context) {
	month, err := getMonthOffset(c)
	if err != nil {
		handleError(c, err, err.Error(), http.StatusBadRequest)
	}

	pagination, err := framework.NewPaginationFromContext(c)
	if err != nil {
		handleError(c, err, "Invalid pagination parameters", http.StatusBadRequest)
	}

	apiKey := helpers.ApiKey(c)
	//if apiKey == "" {
	//	handleError(c, nil, "Cannot get audit data with no api key", http.StatusBadRequest)
	//	return
	//}

	audits, total, err := r.auditRepo.GetByDateAndApiKey(*month, apiKey, pagination.Size, pagination.Offset)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get audit logs"), http.StatusInternalServerError)
		return
	}

	paginator(c, total, *pagination)
	jsonResponse(c, audits)
	c.Header("Cache-Control", "max-age=60")
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
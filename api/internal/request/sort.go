package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
)

var (
	ErrInvalidSortValue = errors.New("invalid sort parameter")
)

type Sort struct {
	Field string
	Asc   bool
}

func newSortFromContext(c *gin.Context) (*Sort, error) {
	sortQuery, exists := c.GetQuery("sort")
	if exists == false {
		return nil, nil
	}

	r := regexp.MustCompile(`^([a-zA-Z\-]*):(asc|desc)$`)

	matches := r.FindStringSubmatch(sortQuery)
	if len(matches) != 3 {
		return nil, ErrInvalidSortValue
	}

	return &Sort{matches[1],matches[2] == "asc"}, nil
}
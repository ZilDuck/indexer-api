package framework

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	defaultSize   uint64 = 10
	defaultOffset uint64 = 1
	maxSize       uint64 = 100
)

var (
	ErrInvalidPaginationParameter = errors.New("invalid pagination parameters")
)

type Pagination struct {
	Size   uint64
	Offset uint64
}

func newPagination(size, offset uint64) *Pagination {
	return &Pagination{Size: size, Offset: offset}
}

func NewPaginationFromContext(c *gin.Context) (*Pagination, error) {
	size := defaultSize
	sizeParam, exists := c.GetQuery("limit")
	if exists == true {
		s, err := strconv.ParseUint(sizeParam, 10, 64)
		if err != nil || s == 0 {
			return nil, ErrInvalidPaginationParameter
		}
		if s > maxSize {
			s = maxSize
		}
		size = s
	}

	offset := defaultOffset
	offsetParam, exists := c.GetQuery("offset")
	if exists == true {
		o, err := strconv.ParseUint(offsetParam, 10, 64)
		if err != nil {
			return nil, ErrInvalidPaginationParameter
		}
		offset = o
	}

	return newPagination(size, offset), nil
}
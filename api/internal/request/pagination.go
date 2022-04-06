package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrInvalidPaginationParameter = errors.New("invalid pagination parameters")
)

type Pagination struct {
	Size   int
	Page   int
	Offset int
}

const (
	PaginationDefaultSize int = 10
	PaginationDefaultPage int = 1
	PaginationMaxSize     int = 100
)

func newPaginationFromContext(c *gin.Context) (pagination *Pagination, err error) {
	size := PaginationDefaultSize
	sizeParam, exists := c.GetQuery("size")
	if exists == true {
		size, err = strconv.Atoi(sizeParam)
		if err != nil || size == 0 {
			return nil, ErrInvalidPaginationParameter
		}
		if size > PaginationMaxSize {
			size = PaginationMaxSize
		}
	}

	page := PaginationDefaultPage
	pageParam, exists := c.GetQuery("page")
	if exists == true {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			return nil, ErrInvalidPaginationParameter
		}
		if page == 0 {
			page = 1
		}
	}

	return &Pagination{Size: size, Page: page, Offset: (page * size) - size}, nil
}

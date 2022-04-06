package request

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Request struct {
	context *gin.Context
	errors  []error

	*Pagination
	*Sort
	*Parameters
}

func NewRequest(c *gin.Context) Request {
	return Request{context: c}
}

func NewPaginatedRequest(c *gin.Context) Request {
	return Request{context: c}.Paginated()
}

func (r Request) Paginated() Request {
	if pagination, err := newPaginationFromContext(r.context); err == nil {
		r.Pagination = pagination
	} else {
		r.errors = append(r.errors, err)
	}

	return r
}

func (r Request) Sortable() Request {
	if sort, err := newSortFromContext(r.context); err == nil {
		r.Sort = sort
	} else {
		r.errors = append(r.errors, err)
	}

	return r
}

func (r Request) WithParams(params Parameters) Request {
	for idx := range params {
		val := r.context.DefaultQuery(params[idx].Name, params[idx].DefaultValue)

		if params[idx].List {
			vals := strings.Split(val, params[idx].ListSeparator)
			params[idx].Value = vals
		}

		switch params[idx].Type {
		case INT:
			value, err := strconv.Atoi(val)
			if err != nil {
				r.errors = append(r.errors, errors.New(fmt.Sprintf("Invalid query parameter %s (%s)", params[idx].Name, val)))
				continue
			}
			if allowed := params[idx].IsAllowed(value); !allowed {
				r.errors = append(r.errors, errors.New(fmt.Sprintf("%s is not a support value of %s", val, params[idx].Name)))
				continue
			}
			params[idx].Value = value

		case INTLIST:
			values := make([]int, 0)
			for _, val := range strings.Split(val, params[idx].ListSeparator) {
				if val == "" {
					continue
				}
				value, err := strconv.Atoi(val)
				if err != nil {
					r.errors = append(r.errors, errors.New(fmt.Sprintf("Invalid query parameter %s (%s)", params[idx].Name, val)))
					continue
				}
				if allowed := params[idx].IsAllowed(value); !allowed {
					r.errors = append(r.errors, errors.New(fmt.Sprintf("%s is not a support value of %s", val, params[idx].Name)))
					continue
				}
				values = append(values, value)
			}
			params[idx].Value = values

		case STRING:
			if allowed := params[idx].IsAllowed(val); !allowed {
				r.errors = append(r.errors, errors.New(fmt.Sprintf("%s is not a support value of %s", val, params[idx].Name)))
				continue
			}
			params[idx].Value = val

		case STRINGLIST:
			values := make([]string, 0)
			for _, val := range strings.Split(val, params[idx].ListSeparator) {
				if val == "" {
					continue
				}
				if allowed := params[idx].IsAllowed(val); !allowed {
					r.errors = append(r.errors, errors.New(fmt.Sprintf("%s is not a support value of %s", val, params[idx].Name)))
					continue
				}
				values = append(values, val)
			}
			params[idx].Value = values
		}
	}

	r.Parameters = &params

	return r
}

func (r Request) Errors() []error {
	return r.errors
}
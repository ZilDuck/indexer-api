package request

import (
	"github.com/gin-gonic/gin"
)

func GetAllContractsRequest(c *gin.Context) Request {
	params := []Parameter{{
		Name:          "from",
		Type:          INT,
		AllowedValues: nil,
		DefaultValue:  "0",
	},{
		Name:          "shape",
		Type:          STRINGLIST,
		List:          true,
		ListSeparator: ",",
		AllowedValues: []interface{}{"ZRC1", "ZRC2", "ZRC3", "ZRC4", "ZRC6"},
	}}

	return NewRequest(c).
		Paginated().
		Sortable().
		WithParams(params)
}


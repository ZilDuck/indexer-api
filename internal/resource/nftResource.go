package resource

import (
	"fmt"
	"github.com/dantudor/zilkroad-txapi/internal/service"
	"github.com/gin-gonic/gin"
)

type NftResource interface {
	GetNftsOwnedByAddress(c *gin.Context)
}

type nftResource struct {
	nftService service.NFTService
}

func NewNftResource(nftService service.NFTService) NftResource {
	return nftResource{nftService}
}

func (r nftResource) GetNftsOwnedByAddress(c *gin.Context) {
	nfts, _, err := r.nftService.GetForAddress(c.Param("ownerAddr"), 0, 500)
	if err != nil {
		errorInternalServerError(c, fmt.Sprintf("Failed to get nfts for address: %s", c.Param("ownerAddr")))
		return
	}

	c.JSON(200, nfts)
}

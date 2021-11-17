package resource

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/service"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"strings"
)

type NftResource interface {
	GetNftsOwnedByAddress(c *gin.Context)
}

type nftResource struct {
	nftService service.NFTService
	cache      persistence.CacheStore
}

func NewNftResource(nftService service.NFTService, cache persistence.CacheStore) NftResource {
	return nftResource{nftService, cache}
}

func (r nftResource) GetNftsOwnedByAddress(c *gin.Context) {
	ownerAddr := strings.ToLower(c.Param("ownerAddr"))

	nfts, _, err := r.nftService.GetForAddress(ownerAddr, 0, 10000)
	if err != nil {
		errorInternalServerError(c, fmt.Sprintf("Failed to get nfts for address: %s", ownerAddr))
		return
	}

	c.Header("Cache-Miss", "true")

	c.JSON(200, nfts)
}

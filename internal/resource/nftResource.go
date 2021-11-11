package resource

import (
	"fmt"
	"github.com/dantudor/zilkroad-txapi/internal/service"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"log"
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
	var data interface{}
	key := cache.CreateKey(c.Request.URL.RequestURI())
	log.Println(key)
	if err := r.cache.Get(key, &data); err != nil {
		log.Println(err.Error())
	}
	log.Println(data)
	nfts, _, err := r.nftService.GetForAddress(c.Param("ownerAddr"), 0, 10000)
	if err != nil {
		errorInternalServerError(c, fmt.Sprintf("Failed to get nfts for address: %s", c.Param("ownerAddr")))
		return
	}

	c.Header("Cache-Miss", "true")

	c.JSON(200, nfts)
}

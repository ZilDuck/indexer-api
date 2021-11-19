package resource

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/factory"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type NftResource struct {
	nftRepo repository.NftRepository
}

func NewNftResource(nftRepo repository.NftRepository) NftResource {
	return NftResource{nftRepo}
}

func (r NftResource) GetNftsOwnedByAddress(c *gin.Context) {
	ownerAddr := strings.ToLower(c.Param("ownerAddr"))

	nfts, total, err := r.nftRepo.GetForAddress(ownerAddr, 0, 1)
	if err != nil {
		msg := fmt.Sprintf("Failed to get nfts for address: %s", ownerAddr)

		zap.L().With(zap.Error(err)).Error(msg)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": msg, "status": http.StatusInternalServerError},
		)

		return
	}
	zap.S().Infof("Found %d NFT For %s", total, ownerAddr)

	c.Header("Cache-Control", "max-age=60")

	c.JSON(200, factory.NftsIndexToDto(nfts))
}

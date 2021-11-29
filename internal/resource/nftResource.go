package resource

import (
	"fmt"
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

	contracts, err := r.nftRepo.GetForAddress(network(c), ownerAddr)
	if err != nil {
		msg := fmt.Sprintf("Failed to get nfts for address: %s", ownerAddr)

		zap.L().With(zap.Error(err)).Error(msg)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": msg, "status": http.StatusInternalServerError},
		)

		return
	}

	c.Header("Cache-Control", "max-age=60")

	c.JSON(200, contracts)
}

func (r NftResource) GetContractNfts(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	nfts, total, err := r.nftRepo.GetForContract(network(c), contractAddr, 10000, 1)
	if err != nil {
		msg := fmt.Sprintf("Failed to get nfts for contract: %s", contractAddr)

		zap.L().With(zap.Error(err)).Error(msg)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": msg, "status": http.StatusInternalServerError},
		)

		return
	}
	zap.S().Infof("Found %d NFT For %s", total, contractAddr)

	c.Header("Cache-Control", "max-age=60")

	c.JSON(200, nfts)
}

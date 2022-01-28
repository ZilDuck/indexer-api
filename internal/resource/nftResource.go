package resource

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/mapper"
	"github.com/ZilDuck/indexer-api/internal/metadata"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

type NftResource struct {
	nftRepo      repository.NftRepository
	metadata     metadata.Service
}

func NewNftResource(nftRepo repository.NftRepository, metadata metadata.Service) NftResource {
	return NftResource{nftRepo, metadata}
}

func (r NftResource) GetContractNfts(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	nfts, _, err := r.nftRepo.GetForContract(network(c), contractAddr, 10000, 1)
	if err != nil {
		msg := fmt.Sprintf("Failed to get nfts for contract: %s", contractAddr)
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": msg, "status": http.StatusInternalServerError})
		return
	}

	c.Header("Cache-Control", "max-age=60")
	c.JSON(200, mapper.NftEntitysToDtos(nfts))
}

func (r NftResource) GetContractNft(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))
	tokenId, err := strconv.ParseUint(c.Param("tokenId"), 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Invalid token id: %s", c.Param("tokenId"))
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": msg, "status": http.StatusBadRequest})
		return
	}

	nft, err := r.nftRepo.GetForContractByTokenId(network(c), contractAddr, tokenId)
	if err != nil {
		msg := fmt.Sprintf("Failed to get %d nft for contract: %s", tokenId, contractAddr)
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": msg, "status": http.StatusInternalServerError})
		return
	}

	c.Header("Cache-Control", "max-age=60")
	c.JSON(200, mapper.NftEntityToDto(*nft))
}

func (r NftResource) GetContractNftMetadata(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))
	tokenId, err := strconv.ParseUint(c.Param("tokenId"), 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Invalid token id: %s", c.Param("tokenId"))
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": msg, "status": http.StatusBadRequest})
		return
	}

	nft, err := r.nftRepo.GetForContractByTokenId(network(c), contractAddr, tokenId)
	if err != nil {
		msg := fmt.Sprintf("Failed to get %d nft of contract: %s", tokenId, contractAddr)
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": msg, "status": http.StatusInternalServerError})
		return
	}

	md, err := r.metadata.GetZrc6Metadata(*nft)
	if err != nil {
		msg := fmt.Sprintf("Failed to get metadata for %d nft of contract: %s", tokenId, contractAddr)
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": msg, "status": http.StatusInternalServerError})
		return
	}

	c.Header("Cache-Control", "max-age=60")
	c.JSON(200, md)
}

func (r NftResource) GetContractNftAsset(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))
	tokenId, err := strconv.ParseUint(c.Param("tokenId"), 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Invalid token id: %s", c.Param("tokenId"))
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	nft, err := r.nftRepo.GetForContractByTokenId(network(c), contractAddr, tokenId)
	if err != nil {
		msg := fmt.Sprintf("Failed to get %d nft of contract: %s", tokenId, contractAddr)
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if nft.MediaUri == "" {
		msg := fmt.Sprintf("Asset not found: %s", c.Param("tokenId"))
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	media, contentType, err := r.metadata.GetZrc6Media(*nft)
	if err != nil {
		zap.L().With(zap.Error(err)).Error("failed to get zrc6 media")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, media); err != nil {
		zap.L().With(zap.Error(err)).Error("failed to encode zrc6 media")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	zap.L().With(zap.String("contractAddr", contractAddr), zap.Uint64("tokenId", tokenId)).Info("Serving binary content")

	c.Header("Cache-Control", "max-age=60")
	c.Data(http.StatusOK, contentType, buf.Bytes())
}

func (r NftResource) GetNftsOwnedByAddress(c *gin.Context) {
	ownerAddr := strings.ToLower(c.Param("ownerAddr"))

	contracts, err := r.nftRepo.GetForAddress(network(c), ownerAddr)
	if err != nil {
		msg := fmt.Sprintf("Failed to get nfts for address: %s", ownerAddr)
		zap.L().With(zap.Error(err)).Error(msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": msg, "status": http.StatusInternalServerError})
		return
	}

	c.Header("Cache-Control", "max-age=60")
	c.JSON(200, contracts)
}

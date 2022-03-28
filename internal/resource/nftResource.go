package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/framework"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/ZilDuck/indexer-api/internal/mapper"
	"github.com/ZilDuck/indexer-api/internal/messenger"
	"github.com/ZilDuck/indexer-api/internal/metadata"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type NftResource struct {
	nftRepo        repository.NftRepository
	actionRepo     repository.NftActionRepository
	messageService messenger.MessageService
	metadata       metadata.Service
}

func NewNftResource(
	nftRepo repository.NftRepository,
	actionRepo repository.NftActionRepository,
	messageService messenger.MessageService,
	metadata metadata.Service,
) NftResource {
	return NftResource{nftRepo, actionRepo, messageService, metadata}
}

func (r NftResource) GetContractNfts(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))
	pagination, err := framework.NewPaginationFromContext(c)
	if err != nil {
		handleError(c, err, "Invalid pagination parameters", http.StatusBadRequest)
	}

	nfts, total, err := r.nftRepo.GetForContract(helpers.Network(c), contractAddr, pagination.Size, pagination.Offset)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get nfts for contract: %s", contractAddr), http.StatusInternalServerError)
		return
	}

	paginator(c, total, *pagination)

	c.Header("Cache-Control", "no-cache")
	c.JSON(200, mapper.NftToDtos(nfts))
}

func (r NftResource) GetContractNft(c *gin.Context) {
	contractAddr, tokenId, err := r.getContractAndTokenId(c)
	if err != nil {
		return
	}

	nft, err := r.nftRepo.GetForContractByTokenId(helpers.Network(c), *contractAddr, *tokenId)
	if err != nil {
		if errors.Is(err, repository.ErrNftNotFound) {
			handleError(c, err, "NFT not found", http.StatusNotFound)
		} else {
			handleError(c, err, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	c.Header("Cache-Control", "no-cache")
	c.JSON(200, mapper.NftToDto(*nft))
}

func (r NftResource) GetContractNftMetadata(c *gin.Context) {
	contractAddr, tokenId, err := r.getContractAndTokenId(c)
	if err != nil {
		return
	}

	nft, err := r.nftRepo.GetForContractByTokenId(helpers.Network(c), *contractAddr, *tokenId)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get %d nft of contract: %s", tokenId, *contractAddr), http.StatusInternalServerError)
		return
	}

	md, err := r.metadata.GetZrc6Metadata(*nft)
	if err != nil {
		msg := fmt.Sprintf("Failed to get metadata for %d nft of contract: %s", tokenId, *contractAddr)
		handleError(c, err, msg, http.StatusInternalServerError)
		return
	}

	c.Header("Cache-Control", "no-cache")
	c.JSON(200, md)
}

func (r NftResource) GetContractNftActions(c *gin.Context) {
	contractAddr, tokenId, err := r.getContractAndTokenId(c)
	if err != nil {
		return
	}

	actions, _, err := r.actionRepo.GetByContractAndTokenId(helpers.Network(c), *contractAddr, *tokenId)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get %d nft of contract: %s", *tokenId, *contractAddr), http.StatusInternalServerError)
		return
	}

	c.JSON(200, mapper.ActionsToDtos(actions))
}

func (r NftResource) GetNftsOwnedByAddress(c *gin.Context) {
	ownerAddr := getAddress(c.Param("ownerAddr"))
	shape := strings.ToUpper(c.DefaultQuery("shape", ""))

	nfts, err := r.nftRepo.GetForAddress(helpers.Network(c), ownerAddr, shape)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get nfts for address: %s", ownerAddr), http.StatusInternalServerError)
		return
	}

	c.JSON(200, nfts)
}

func (r NftResource) RefreshMetadata(c *gin.Context) {
	contractAddr, tokenId, err := r.getContractAndTokenId(c)
	if err != nil {
		return
	}

	message := messenger.RefreshMetadata{
		Contract: *contractAddr,
		TokenId: *tokenId,
	}

	msgJson, _ := json.Marshal(message)
	if err := r.messageService.SendMessage(helpers.Network(c), messenger.MetadataRefresh, msgJson); err != nil {
		handleError(c, err, "Failed to queue metadata refresh", http.StatusBadRequest)
	}
	message.Sent = true

	c.JSON(200, message)
	c.Header("Cache-Control", "no-store")
}

func (r NftResource) getContractAndTokenId(c *gin.Context) (*string, *uint64, error) {
	contractAddr := getAddress(c.Param("contractAddr"))
	tokenId, err := strconv.ParseUint(c.Param("tokenId"), 0, 64)

	if err != nil {
		msg := fmt.Sprintf("Invalid token id: %s", c.Param("tokenId"))
		handleError(c, err, msg, http.StatusBadRequest)
		return nil, nil, errors.New(msg)
	}

	return &contractAddr, &tokenId, nil
}

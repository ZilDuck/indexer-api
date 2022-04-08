package resource

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/dev"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/ZilDuck/indexer-api/internal/mapper"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/ZilDuck/indexer-api/internal/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ContractResource struct {
	contractRepo repository.ContactRepository
	contractStateRepo repository.ContactStateRepository
	nftRepo      repository.NftRepository
}

func NewContractResource(contractRepo repository.ContactRepository, contractStateRepo repository.ContactStateRepository, nftRepo repository.NftRepository) ContractResource {
	return ContractResource{contractRepo, contractStateRepo,nftRepo}
}

func (r ContractResource) GetContract(c *gin.Context) {
	contractAddr := getAddress(c.Param("contractAddr"))

	contract, err := r.contractRepo.GetContract(helpers.Network(c), contractAddr)
	if err != nil {
		if errors.Is(err, repository.ErrContractNotFound) {
			handleError(c, err, "Contract not found", http.StatusNotFound)
		} else {
			handleError(c, err, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	jsonResponse(c, mapper.ContractToDto(*contract))
}

func (r ContractResource) GetContracts(c *gin.Context) {
	req := request.GetAllContractsRequest(c)
	if req.Errors() != nil {
		handleError(c, req.Errors()[0], "Invalid request", http.StatusBadRequest)
		return
	}

	dev.Dump(req)

	contracts, total, err := r.contractRepo.GetAll(
		helpers.Network(c),
		req.Pagination,
		req.Sort,
		req.Parameters.Uint64("from"),
		req.Parameters.StringList("shape"))
	if err != nil {
		handleError(c, err, "Internal server error", http.StatusInternalServerError)
		return
	}

	paginator(c, total, req.Pagination)
	jsonResponse(c, mapper.ContractsToDtos(contracts))
}

func (r ContractResource) GetCode(c *gin.Context) {
	contractAddr := getAddress(c.Param("contractAddr"))

	contract, err := r.contractRepo.GetContract(helpers.Network(c), contractAddr)
	if err != nil {
		if errors.Is(err, repository.ErrContractNotFound) {
			handleError(c, err, "Contract not found", http.StatusNotFound)
		} else {
			handleError(c, err, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(contract.Code))
}

func (r ContractResource) GetAttributes(c *gin.Context) {
	contractAddr := getAddress(c.Param("contractAddr"))
	tokenIds := getTokenIdsFromQueryList(c.Query("tokenIds"))

	attributes, err := r.nftRepo.GetForContractAttributes(helpers.Network(c), contractAddr, tokenIds)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get attributes: %s", contractAddr), 500)
		return
	}

	jsonResponse(c, attributes)
}

func (r ContractResource) GetState(c *gin.Context) {
	contractAddr := getAddress(c.Param("contractAddr"))
	filters := strings.Split(c.Query("filters"), ",")

	state, err := r.contractStateRepo.GetState(helpers.Network(c), contractAddr)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get state: %s", contractAddr), 500)
		return
	}

	jsonResponse(c, mapper.StateToDto(*state, filters))
}

func (r ContractResource) GetContractsOwnedByAddress(c *gin.Context) {
	ownerAddr := getAddress(c.Param("ownerAddr"))
	details := strings.ToLower(c.DefaultQuery("details", "false"))

	contractAddrs, err := r.contractStateRepo.GetAllAddressesOwnedBy(helpers.Network(c), ownerAddr)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get contracts for address: %s", ownerAddr), http.StatusInternalServerError)
		return
	}

	if details == "false" {
		jsonResponse(c, contractAddrs)
		return
	}

	contracts, err := r.contractRepo.GetContracts(helpers.Network(c), contractAddrs...)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get contracts for address: %s", ownerAddr), http.StatusInternalServerError)
		return
	}
	jsonResponse(c, mapper.ContractsToDtos(contracts))
}

func getTokenIdsFromQueryList(query string) (tokenIds []uint64) {
	for _, el := range strings.Split(query, ",") {
		if tokenId, err := strconv.ParseUint(el, 0, 64); err == nil {
			if tokenId != 0 {
				tokenIds = append(tokenIds, tokenId)
			}
		}
	}
	return
}
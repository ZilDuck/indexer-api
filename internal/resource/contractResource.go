package resource

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/helpers"
	"github.com/ZilDuck/indexer-api/internal/mapper"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
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
	contractAddr := strings.ToLower(c.Param("contractAddr"))

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

func (r ContractResource) GetCode(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

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
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	attributes, err := r.nftRepo.GetForContractAttributes(helpers.Network(c), contractAddr)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get attributes: %s", contractAddr), 500)
		return
	}

	jsonResponse(c, attributes)
}

func (r ContractResource) GetState(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	state, err := r.contractStateRepo.GetState(helpers.Network(c), contractAddr)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get state: %s", contractAddr), 500)
		return
	}

	jsonResponse(c, state)
}

func (r ContractResource) GetContractsOwnedByAddress(c *gin.Context) {
	ownerAddr := strings.ToLower(c.Param("ownerAddr"))
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
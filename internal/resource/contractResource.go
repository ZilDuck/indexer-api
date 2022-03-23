package resource

import (
	"errors"
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/mapper"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ContractResource struct {
	contractRepo repository.ContactRepository
	nftRepo      repository.NftRepository
}

func NewContractResource(contractRepo repository.ContactRepository, nftRepo repository.NftRepository) ContractResource {
	return ContractResource{contractRepo, nftRepo}
}

func (r ContractResource) GetContract(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	contract, err := r.contractRepo.GetContract(network(c), contractAddr)
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

	contract, err := r.contractRepo.GetContract(network(c), contractAddr)
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

	attributes, err := r.nftRepo.GetForContractAttributes(network(c), contractAddr)
	if err != nil {
		handleError(c, err, fmt.Sprintf("Failed to get attributes: %s", contractAddr), 500)
		return
	}

	jsonResponse(c, attributes)
	c.Header("Cache-Control", "max-age=60")
}

func (r ContractResource) GetContractsOwnedByAddress(c *gin.Context) {
	ownerAddr := strings.ToLower(c.Param("ownerAddr"))
	details := strings.ToLower(c.DefaultQuery("details", "false"))

	if details == "true" {
		contracts, err := r.contractRepo.GetAllOwnedBy(network(c), ownerAddr)
		if err != nil {
			msg := fmt.Sprintf("Failed to get contracts for address: %s", ownerAddr)
			handleError(c, err, msg, http.StatusInternalServerError)
			return
		}
		jsonResponse(c, mapper.ContractsToDtos(contracts))
	} else {
		contractAddrs, err := r.contractRepo.GetAllAddressesOwnedBy(network(c), ownerAddr)
		if err != nil {
			msg := fmt.Sprintf("Failed to get contracts for address: %s", ownerAddr)
			handleError(c, err, msg, http.StatusInternalServerError)
			return
		}
		jsonResponse(c, contractAddrs)
	}
	c.Header("Cache-Control", "max-age=60")
}
package resource

import (
	"fmt"
	"github.com/ZilDuck/indexer-api/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type ContractResource struct {
	contactRepo repository.ContactRepository
}

func NewContractResource(contactRepo repository.ContactRepository) ContractResource {
	return ContractResource{contactRepo}
}

func (r ContractResource) GetContract(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	contract, err := r.contactRepo.GetContract(network(c), contractAddr)
	if err != nil {
		msg := fmt.Sprintf("Failed to get contract: %s", contractAddr)

		zap.L().With(zap.Error(err)).Error(msg)

		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": msg, "status": http.StatusInternalServerError},
		)

		return
	}

	c.Header("Cache-Control", "max-age=60")

	c.JSON(200, contract)
}

func (r ContractResource) GetContracts(c *gin.Context) {
	contracts, _, err := r.contactRepo.GetContracts(network(c))
	if err != nil {
		msg := fmt.Sprintf("Failed to get contracts")

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

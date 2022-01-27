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
	contractRepo repository.ContactRepository
}

func NewContractResource(contractRepo repository.ContactRepository) ContractResource {
	return ContractResource{contractRepo}
}

func (r ContractResource) GetContract(c *gin.Context) {
	contractAddr := strings.ToLower(c.Param("contractAddr"))

	contract, err := r.contractRepo.GetContract(network(c), contractAddr)
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
package resource

import (
	"github.com/ZilDuck/indexer-api/internal/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AdminResource struct {
	authService auth.Service
}

type CreateClient struct {
	Name string `json:"username"`
}

func NewAdminResource(authService auth.Service) AdminResource {
	return AdminResource{authService}
}

func (r AdminResource) GetClients(c *gin.Context) {
	jsonResponse(c, auth.GetApiClients())
}

func (r AdminResource) CreateClient(c *gin.Context) {
	body := CreateClient{}
	if err := c.BindJSON(&body); err != nil {
		handleError(c, err, "Failed to create new client", http.StatusBadRequest)
		return
	}

	client, err := r.authService.CreateClient(body.Name, true)
	if err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(c, client)
}

func (r AdminResource) DisableClient(c *gin.Context) {
	client, err := auth.GetClientByUsername(c.Param("username"))
	if err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}
	zap.S().Infof("Disable client %s", client.Username)

	client.Active = false

	if err := r.authService.UpdateClient(*client); err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "text/plain", []byte(""))
}

func (r AdminResource) EnableClient(c *gin.Context) {
	client, err := auth.GetClientByUsername(c.Param("username"))
	if err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	client.Active = true

	if err := r.authService.UpdateClient(*client); err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "text/plain", []byte(""))
}

func (r AdminResource) DeleteClient(c *gin.Context) {
	username := c.Param("username")

	client, err := auth.GetClientByUsername(username)
	if err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.authService.DeleteClient(*client); err != nil {
		handleError(c, err, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "text/plain", []byte(""))
}

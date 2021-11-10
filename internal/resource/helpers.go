package resource

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorNetworkNotAvailable(c *gin.Context) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		gin.H{"message": "Network not available", "status": http.StatusNotFound},
	)
}

func errorNotFound(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		gin.H{"message": msg, "status": http.StatusNotFound},
	)
}

func errorBadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		gin.H{"message": msg, "status": http.StatusBadRequest},
	)
}

func errorRequestError(c *gin.Context, err error) {
	errorInternalServerError(c, "Failed to process request:"+err.Error())
}

func errorInternalServerError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		gin.H{"message": msg, "status": http.StatusInternalServerError})
}

func handleError(c *gin.Context, err error, status int) {
	c.AbortWithStatusJSON(status, gin.H{
		"status":  status,
		"message": err.Error(),
	})
}

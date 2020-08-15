package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iino123/golang-microservices/src/api/domain/repositories"
	"github.com/iino123/golang-microservices/src/api/services"
	"github.com/iino123/golang-microservices/src/api/utils/errors"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	clientId := c.GetHeader("X-Client-Id")

	result, err := services.RepositoryService.CreateRepo(clientId, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

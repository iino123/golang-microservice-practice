package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iino123/golang-microservice-practice/mvc/services"
	"github.com/iino123/golang-microservice-practice/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	//c.JSON(http.StatusOK, user)
	utils.Respond(c, http.StatusOK, user)
}

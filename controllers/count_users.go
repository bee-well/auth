package controllers

import (
	"net/http"

	"github.com/bee-well/auth/services"
	"github.com/gin-gonic/gin"
)

type countResponse struct {
	Count int `json:"count"`
}

func CountUsers(c *gin.Context) {
	count, err := services.GetUserCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something went wrong, please try again later")
		return
	}
	c.JSON(http.StatusOK, countResponse{count})
}

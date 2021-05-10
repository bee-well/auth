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
	_, err := services.GetUserCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

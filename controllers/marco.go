package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Marco(c *gin.Context) {
	c.String(http.StatusOK, "polo")
}

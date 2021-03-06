package controllers

import (
	"net/http"

	"github.com/bee-well/auth/services"
	"github.com/gin-gonic/gin"
)

type signInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {
	var payload signInPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Could not process the request body.")
		return
	}

	if payload.Email == "" || payload.Password == "" {
		c.JSON(http.StatusBadRequest, "Could not process the request body.")
		return
	}

	if token, err := services.SignIn(payload.Email, payload.Password); err != nil {
		c.JSON(http.StatusForbidden, err.Error())
	} else {
		c.JSON(http.StatusOK, token)
	}
}

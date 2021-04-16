package controllers

import (
	"net/http"
	"strings"

	"github.com/bee-well/auth/services"
	"github.com/gin-gonic/gin"
)

type signUpPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func SignUp(c *gin.Context) {
	var payload signUpPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide sufficient information")
		return
	}

	if !strings.Contains(payload.Email, "@") || !strings.Contains(payload.Email, ".") {
		c.JSON(http.StatusBadRequest, "Invalid email address")
		return
	}

	if len(payload.Password) < 8 {
		c.JSON(http.StatusBadRequest, "Insufficient password length")
		return
	}

	if payload.FirstName == "" || payload.LastName == "" {
		c.JSON(http.StatusBadRequest, "Please provide both a first and last name")
		return
	}

	if err := services.SignUp(
		payload.Email,
		payload.Password,
		payload.FirstName,
		payload.LastName,
	); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "yes!")
}

package controllers

import (
	"net/http"

	"github.com/bee-well/auth/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	tokens := c.Request.Header["Authorization"]
	if len(tokens) == 0 {
		c.JSON(http.StatusForbidden, "no authorization token")
		return
	}

	token := tokens[0]

	apiToken, ok := services.Verify(token)
	if !ok {
		c.JSON(http.StatusForbidden, "bad token")
		return
	}

	user, err := services.GetUser(apiToken.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "an error occurred, please try again later")
		return
	}

	c.JSON(http.StatusOK, user)
}

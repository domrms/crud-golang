package controller

import (
	"crud-golang/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("Invalid user data")
	c.JSON(err.Code, err)
}

func FindUserByEmail(c *gin.Context) {
	err := rest_err.NewBadRequestError("Invalid email data")
	c.JSON(err.Code, err)
}

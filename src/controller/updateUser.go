package controller

import (
	"crud-golang/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func UpdateUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("Invalid user data")
	c.JSON(err.Code, err)
}

package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParseUintParam(c *gin.Context, paramName string) (uint, error) {
	paramValue := c.Param(paramName)
	parsedID, err := strconv.Atoi(paramValue)
	if err != nil {
		return 0, err
	}
	return uint(parsedID), nil
}

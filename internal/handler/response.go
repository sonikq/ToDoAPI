package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorRespose struct {
	Message string `json: "message"`
}

type statusResposne struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context,statusCode int, message string){
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorRespose{message})
}

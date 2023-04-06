package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func sendResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	sendResponse(c, status, response{Data: data})
}

func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Status:  status,
		Code:    strings.ReplaceAll(strings.ToUpper(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
	}

	sendResponse(c, err.Status, err)
}

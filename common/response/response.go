package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string 	`json:"message"`
	Data interface{} `json:"data"`
}

func SendSuccess(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Message: msg,
		Data: data,
	})
	return
}

func SendError(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: code,
		Message: msg,
		Data: data,
	})
	return
}



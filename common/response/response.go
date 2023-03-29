package response

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendSuccess 成功响应
func SendSuccess(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}

// SendError 失败响应
func SendError(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}

//ReturnJsonFromString 将json字符串以标准json格式返回
func ReturnJsonFromString(c *gin.Context, httpCode int, jsonStr string) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(httpCode, jsonStr)
}


// SendFail 失败的业务逻辑
func SendFail(c *gin.Context, dataCode int, msg string, data interface{}) {
	SendError(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}


// ParseParam 解析URL参数
func ParseParam(c *gin.Context, key string) string {
	return c.Param(key)
}

//

// ParseJSON 解析json类型参数到结构体
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.New(fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

// ParseQuery 只绑定查询字符串到结构体
func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.New(fmt.Sprintf("Parse request query failed: %s", err.Error()))
	}
	return nil
}

// ParseForm 绑定form表单数据到结构体
func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.New(fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

func FailWithMessage(message string, c *gin.Context) {
	SendError(c, 500, 999, message, map[string]interface{}{})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	SendError(c, 500, 999, message, data)
}


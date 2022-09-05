package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

// LoggerMiddleware 接收gin框架默认的日志
func LoggerMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		logger := logrus.New()
		startTime := time.Now()
		c.Next()  // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds() / 1000000))))
		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := logger.WithFields(logrus.Fields{
			"SpendTime": spendTime,
			"path": url,
			"Method": method,
			"status": statusCode,
		})
		if len(c.Errors) > 0 {  // 矿建内部错误
			Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode >= 500 {
			Log.Error()
		} else if statusCode >= 400 {
			Log.Warn()
		}else {
			Log.Info()
		}
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover(); err != nil {
			fields := make(map[string]interface{})
			fields["ip"] = c.ClientIP()
			fields["method"] = c.Request.Method
			fields["url"] = c.Request.URL.String()
			fields["proto"] = c.Request.Proto
			fields["header"] = c.Request.Header
			fields["user_agent"] = c.GetHeader("User-Agent")
			fields["content_length"] = c.Request.ContentLength
		}
		c.Abort()
	}
}

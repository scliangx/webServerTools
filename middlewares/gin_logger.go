package middware

import (
	"github.com/gin-gonic/gin"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		p := c.Request.URL.Path
		method := c.Request.Method

		fields := make(map[string]interface{})
		fields["path"] = p
		fields["ip"] = c.ClientIP()
		fields["method"] = method
		fields["url"] = c.Request.URL.String()
		fields["proto"] = c.Request.Proto
		fields["header"] = c.Request.Header
		fields["user_agent"] = c.GetHeader("User-Agent")
		fields["content_length"] = c.Request.ContentLength

		c.Next()

	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := recover();err != nil{
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

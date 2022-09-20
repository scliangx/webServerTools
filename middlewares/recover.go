package middware

import (
	"github.com/gin-gonic/gin"
)

// GinRecovery recover掉项目可能出现的panic
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
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
			// 在此处处理panic
			// ...
			c.Abort()
		}()
		// 放行
		c.Next()
	}
}

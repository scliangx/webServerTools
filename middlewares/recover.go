package middware

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scliang-strive/webServerTools/common/response"
)

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(io.Discard, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即可
		// 这里的 err 数据类型为 ：runtime.boundsError  ，需要转为普通数据类型才可以输出
		response.SendError(c, http.StatusHTTPVersionNotSupported, 0, fmt.Sprintf("%s", err), "")
	})
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/scliang-strive/webServerTools/common/response"
	"github.com/sirupsen/logrus"
)

func ApiRouter(router *gin.Engine) {
	router.GET("/index", func(c *gin.Context) {
		logrus.Info("-----------index------------")
		response.SendSuccess(c, 200, "", "")
	})
	router.GET("/", func(c *gin.Context) {
		response.SendSuccess(c, 200, "", "")
	})
}

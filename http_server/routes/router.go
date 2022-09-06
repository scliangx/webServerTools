package routes

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(router *gin.Engine) {
	server := router.Group("")
	serverRouter(server)
}

func serverRouter(group *gin.RouterGroup) {

}

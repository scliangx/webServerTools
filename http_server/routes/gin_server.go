package routes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scliang-strive/webServerTools/config"
	middleware "github.com/scliang-strive/webServerTools/middlewares"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	if config.GetConfig().Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		router = gin.New()
		// 载入gin的中间件，关键是第二个中间件，我们对它进行了自定义重写，将可能的 panic 异常等，统一使用 zaplog 接管，保证全局日志打印统一
		router.Use(gin.Logger(), middleware.CustomRecovery())
	} else {
		router = gin.Default()
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
	})
	err := router.Run(fmt.Sprintf(":%d", config.GetConfig().WebConfig.Port))
	if err != nil {
		panic(err)
	}
	ApiRouter(router)
	return router
}

func Run(router *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetConfig().WebConfig.Port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	pid := fmt.Sprintf("%d", os.Getpid())
	_, openErr := os.OpenFile("pid", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openErr == nil {
		_ = ioutil.WriteFile("pid", []byte(pid), 0)
	}
}

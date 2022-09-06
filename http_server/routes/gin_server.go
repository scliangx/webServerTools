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
	"os/signal"
	"time"
)

func InitApiRouter() *gin.Engine {
	var router *gin.Engine
	if config.GetConfig().Debug {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
		router = gin.New()
		// 载入gin的中间件，关键是第二个中间件，我们对它进行了自定义重写，将可能的 panic 异常等
		router.Use(middleware.CustomRecovery())
	} else {
		router = gin.Default()
	}
	router.Use(middleware.LoggerMiddleware())
	router.Static("../static","./static")
	//router.Use(middleware.GinLogger())
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

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

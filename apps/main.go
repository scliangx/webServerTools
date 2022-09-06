package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/scliang-strive/webServerTools/common/logger"
	_ "github.com/scliang-strive/webServerTools/common/logger"
	"github.com/scliang-strive/webServerTools/config"
	"github.com/scliang-strive/webServerTools/http_server/routes"
	"github.com/scliang-strive/webServerTools/internal/db"
	"github.com/scliang-strive/webServerTools/internal/redis"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "web-server",
				Usage: "weg server tools",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "config,c",
						Value: "../config/config.yaml",
						Usage: "config path",
					},
				},
				Action: run,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error("[error] : %s", err)
		panic(err)
	}
	return
}

func run(c *cli.Context) error {
	path := c.String("config")
	config.LoadConfigFromYaml(path)
	fmt.Println("------------load config success------------------")
	logger.LogInit()
	fmt.Println("------------log init success------------------")
	conf := config.GetConfig()
	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	err := db.NewConnection(conf.DB)
	if err != nil {
		panic("database init failed.")
	} else {
		fmt.Println("------------database init success------------------")
	}
	//redis.NewRedisClient()
	redis.InitRedisClientPool(conf.Redis)
	fmt.Println("------------redis init success----------------------")
	router := routes.InitApiRouter()
	routes.Run(router)
	fmt.Println("-------------------apps end------------------")
	return nil
}

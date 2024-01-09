package main

import (
	"fmt"
	"github.com/scliangx/webServerTools/http_server/routes"
	"github.com/scliangx/webServerTools/internal/db"
	"github.com/scliangx/webServerTools/internal/redis"
	"os"
	"runtime"

	"github.com/scliangx/webServerTools/common/logger"
	"github.com/scliangx/webServerTools/config"
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
	}
}

func run(c *cli.Context) error {
	path := c.String("config")
	config.LoadConfigFromYaml(path)
	fmt.Println("------------load config success------------------")
	logger.LogInit()
	fmt.Println("------------log init success------------------")
	conf := config.GetConfig()
	// 初始化组件配置
	{
		err := db.NewConnection(conf.DB)
		if err != nil {
			logrus.Fatalln("database init failed.")
		} else {
			fmt.Println("------------database init success------------------")
		}
		//redis.NewRedisClient()
		redis.InitRedisClientPool(conf.Redis)
		fmt.Println("------------redis init success----------------------")
	}

	routes.InitApiRouter()
	fmt.Println("-------------------apps end------------------")
	return nil
}

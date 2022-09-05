package main

import (
	"github.com/gin-gonic/gin"
	"github.com/scliang-strive/webServerTools/config"
	"github.com/scliang-strive/webServerTools/http_server/routes"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"runtime"
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
		panic(err)
	}
	return
}

func run(c *cli.Context) error {
	path := c.String("config")
	config.LoadConfigFromYaml(path)
	conf := config.GetConfig()

	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	router := routes.InitApiRouter()
	routes.Run(router)
	return nil
}

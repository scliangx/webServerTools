package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "web-server",
				Usage:  "weg server tools",
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
	return nil
}

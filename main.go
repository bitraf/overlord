package main

import (
	"os"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/log"
	"github.com/codegangsta/cli"
)

const APP_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "overlord"
	app.Usage = "Project Overlord"
	app.Version = APP_VER
	app.Action = start
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "path/to/config.yml",
			Usage: "configuration file to use",
		},
		cli.StringFlag{
			Name:  "level",
			Value: "info",
			Usage: "logging level to use",
		},
	}

	app.Run(os.Args)
}

func start(ctx *cli.Context) {
	log.Configure(ctx)
	config.Load(ctx)
}

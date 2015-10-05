package main

import (
	"os"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/logging"
	"github.com/bitraf/overlord/rest"
	"github.com/codegangsta/cli"
)

const APP_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "overlord"
	app.Usage = "Project Overlord"
	app.Version = APP_VER
	app.Action = action
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "path/to/config.yml",
			Usage: "configuration file to use",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "debug logging on/off",
		},
	}

	app.Run(os.Args)
}

func action(ctx *cli.Context) {
	logging.Configure(ctx)
	logging.Infof("Starting Overlord v%s", APP_VER)

	config.Load(ctx)

	db := db.New()
	db.Open()

	server := rest.New(&db)
	server.Start()
}

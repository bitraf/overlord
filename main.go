package main

import (
	"os"

	"github.com/bitraf/overlord/cmd"
	"github.com/codegangsta/cli"
)

const APP_VER = "0.0.1"

func init() {
}

func main() {
	app := cli.NewApp()
	app.Name = "overlord"
	app.Usage = "Project Overlord"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdCheck,
		cmd.CmdStart,
	}
	app.Run(os.Args)
}
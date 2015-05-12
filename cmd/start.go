package cmd

import (
    "github.com/codegangsta/cli"
    "github.com/bitraf/overlord/rest"
)

var CmdStart = cli.Command{
    Name:        "start",
    Usage:       "Start the server",
    Description: `Start the server on specified configuration.`,
    Action:      doStart,
    Flags: []cli.Flag{
        cli.StringFlag{"config, c", "custom/conf/app.ini", "Custom configuration file path", ""},
    },
}

func doStart(ctx *cli.Context) {
    println("Starting the server...")
    rest.StartServer()
}

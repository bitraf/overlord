package cmd

import (
    "github.com/codegangsta/cli"
)

var CmdCheck = cli.Command{
    Name:        "check",
    Usage:       "Check for valid configuration",
    Description: `Check for valid configuration by loading the configurations and checking the database`,
    Action:      doCheck,
    Flags: []cli.Flag{
        cli.StringFlag{"config, c", "custom/conf/app.ini", "Custom configuration file path", ""},
    },
}

func doCheck(ctx *cli.Context) {
    println("Check the configuration...")
}

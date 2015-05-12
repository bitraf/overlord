package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/bitraf/overlord/log"
	"github.com/codegangsta/cli"
)

type Server struct {
	Port int
}

type Database struct {
	User     string
	Password string
}

type Config struct {
	Server
	Database
}

var C Config = Config{
	Server{
		Port: 1234,
	},
	Database{
		User:     "sa",
		Password: "password",
	},
}

func Load(ctx *cli.Context) {
	if !ctx.IsSet("config") {
		log.Debug("Config file not set. Using defaults.")
		return
	}

	name := ctx.String("config")
	fields := log.Fields{"file": name}

	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panicf("Could not find config file.", fields)
	}

	log.Debugf("Reading config file.", fields)
	yaml.Unmarshal(data, &C)
}

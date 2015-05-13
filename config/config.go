package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/bitraf/overlord/log"
	"github.com/codegangsta/cli"
)

type Server struct {
	Addr string
}

type Database struct {
	User     string
	Password string
}

type Config struct {
	Version  string
	Server   Server
	Database Database
}

var C Config = Config{
	Server: Server{
		Addr: ":1234",
	},
	Database: Database{
		User:     "sa",
		Password: "password",
	},
}

func Load(ctx *cli.Context) {
	C.Version = ctx.App.Version

	if !ctx.IsSet("config") {
		log.Debug("Config file not set. Using defaults.")
		return
	}

	name := ctx.String("config")
	fields := log.Fields{"file": name}

	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panicf(err.Error(), fields)
	}

	log.Debugf("Reading config file.", fields)
	yaml.Unmarshal(data, &C)
}

package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/bitraf/overlord/logging"
	"github.com/codegangsta/cli"
)

type Server struct {
	Addr string
}

type Database struct {
	Driver  string
	Url     string
	ShowSQL bool
	Pool    Pool
}

type Pool struct {
	MaxIdle int
	MaxOpen int
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
		Driver:  "postgres",
		Url:     "postgres://postgres:postgres@192.168.99.100:5432/postgres?sslmode=disable",
		ShowSQL: true,
		Pool: Pool{
			MaxIdle: 10,
			MaxOpen: 100,
		},
	},
}

func Load(ctx *cli.Context) {
	C.Version = ctx.App.Version

	if !ctx.IsSet("config") {
		logging.Info("Config file not set. Using defaults.")
		return
	}

	name := ctx.String("config")

	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panicf(err.Error())
	}

	logging.Infof("Reading config file [%s].", name)
	yaml.Unmarshal(data, &C)
}

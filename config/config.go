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
	Name     string
	Addr     string
	User     string
	Password string
	Options  string
	ShowSQL  bool
	Pool     Pool
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
		Name:     "postgres",
		Addr:     "192.168.59.103:5432",
		User:     "p2k12",
		Password: "password",
		Options:  "sslmode=disable",
		ShowSQL:  true,
		Pool: Pool{
			MaxIdle: 10,
			MaxOpen: 100,
		},
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

# Project Overlord

Member database and central API for hackerspace management built with [Go](https://golang.org/).

## Building

Install go (https://golang.org/) and do the following:

    go build

Start using...

    ./overlord

Show help for options...

    ./overlord --help

## Database backend

The database schema is based on the existing p2k12 database to ease the migration.

## Libraries to use

* Command line parser (https://github.com/codegangsta/cli).
* Yaml for configuration (https://github.com/go-yaml/yaml).
* Yaml for configuration (https://github.com/olebedev/config).
* Configuration library (https://github.com/spf13/viper).
* Echo for rest services (https://github.com/labstack/echo).
* Gorm for database to object mapping (https://github.com/jinzhu/gorm).
* Unify the logging into logrus? (https://github.com/Sirupsen/logrus).

## TODO

* Configure database driver using the loaded settings.

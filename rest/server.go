package rest

import (
	"time"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/log"
	"github.com/labstack/echo"
)

type Server struct {
	startTime time.Time
	db        *db.Database
}

func New(db *db.Database) Server {
	return Server{
		db: db,
	}
}

func (server *Server) Start() {
	e := echo.New()
	e.Use(Logger)

	e.Get("/status", server.getStatus)
	e.Get("/checkins", server.getCheckins)
	e.Get("/events", server.getEvents)
	e.Get("/auths", server.getAuthEntries)

	addr := config.C.Server.Addr
	log.Infof("Starting rest endpoins", log.Fields{"addr": addr})

	server.startTime = time.Now()
	e.Run(addr)
}

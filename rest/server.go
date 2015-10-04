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

	addGet(e, "/status", server.getStatus)
	addGet(e, "/checkins", server.getCheckins)
	addGet(e, "/events", server.getEvents)
	addGet(e, "/auths", server.getAuthEntries)
	addGet(e, "/users", server.getUsers)
	addGet(e, "/users/:id", server.getUserById)
	addGet(e, "/members", server.getMembers)
	addGet(e, "/members/:id", server.getMemberById)

	addr := config.C.Server.Addr
	log.Infof("Starting rest endpoins", log.Fields{"addr": addr})

	server.startTime = time.Now()
	e.Run(addr)
}

func addGet(e *echo.Echo, path string, handler echo.Handler) {
	e.Get(path, handler)
	e.Head(path, handler)
}

package rest

import (
	"net/http"
	"time"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/logging"
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

func (s *Server) Start() {
	router := s.NewRouter()

	addr := config.C.Server.Addr
	logging.Infof("Starting server listening on %s", addr)

	s.startTime = time.Now()
	// log.Fatal(http.ListenAndServe(addr, router))
	http.ListenAndServe(addr, router)
}

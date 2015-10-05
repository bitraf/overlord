package rest

import (
	"net/http"
	"time"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/db"
	"github.com/bitraf/overlord/log"
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
	log.Infof("Starting rest endpoins", log.Fields{"addr": addr})

	s.startTime = time.Now()
	// log.Fatal(http.ListenAndServe(addr, router))
	http.ListenAndServe(addr, router)
}

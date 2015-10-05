package rest

import (
	"net/http"
	"time"

	"github.com/bitraf/overlord/config"
	"github.com/bitraf/overlord/model"
)

func (server *Server) getStatus(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	duration := now.Sub(server.startTime).String()

	res := model.Status{
		Version: config.C.Version,
		Uptime:  duration,
	}

	serveJson(w, http.StatusOK, res)
}

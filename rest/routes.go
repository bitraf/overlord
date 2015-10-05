package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", s.getStatus).Methods("GET", "HEAD")

	router.HandleFunc("/users", s.getUsers).Methods("GET", "HEAD")
	router.HandleFunc("/users/{id}", s.getUserById).Methods("GET", "HEAD")

	router.HandleFunc("/members", s.getMembers).Methods("GET", "HEAD")
	router.HandleFunc("/members/{id}", s.getMemberById).Methods("GET", "HEAD")

	router.HandleFunc("/checkins", s.getCheckins).Methods("GET", "HEAD")
	router.HandleFunc("/events", s.getEvents).Methods("GET", "HEAD")
	router.HandleFunc("/auths", s.getAuthEntries).Methods("GET", "HEAD")

	router.HandleFunc("/", s.handleDefault)

	return router
}

func (s *Server) handleDefault(w http.ResponseWriter, r *http.Request) {
	errorJson(w, http.StatusNotFound, "Not found")
}

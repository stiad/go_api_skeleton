package app

import (
	"net/http"
)

func (s *Server) handleHelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.LogDebug("handlers.handleFindOrder")
		s.respond(w, struct {
			Message string `json:"msg"`
		}{"Hello World"}, http.StatusOK)
	}
}

func (s *Server) handleProtectedHelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.LogDebug("handlers.handleFindOrder")
		s.respond(w, struct {
			Message string `json:"msg"`
		}{"Hello World"}, http.StatusOK)
	}
}

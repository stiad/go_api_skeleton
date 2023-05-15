package app

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func (s *Server) middlewareJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json; charset=UTF-8")
		next(w, r)
	}
}

func (s *Server) middlewareApiKeyAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.LogDebug("Testing against Api-Key:", os.Getenv("API_KEY"), "with request using Api-Key:", r.Header.Get("Api-Key"))
		if len(os.Getenv("API_KEY")) > 2 && r.Header.Get("Api-Key") == os.Getenv("API_KEY") {
			next(w, r)
			return
		}
		s.error(w, errors.New("api authentication failed"), http.StatusUnauthorized)
	}
}

func (s *Server) middlewareAccessLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next(w, r)
	}
}

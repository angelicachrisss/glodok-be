package http

import (
	"net/http"

	"glodok-be/pkg/grace"

	"github.com/rs/cors"
)

// Handler ...
type Handler interface {
	GetGlodok(w http.ResponseWriter, r *http.Request)
	InsertGlodok(w http.ResponseWriter, r *http.Request)
	DeleteGlodok(w http.ResponseWriter, r *http.Request)
	UpdateGlodok(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Glodok Handler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}

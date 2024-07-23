package http

import (
	"errors"
	"log"
	"net/http"

	"glodok-be/pkg/response"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	// Jika tidak ditemukan, jangan diubah.
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Health Check
	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	// Tambahan Prefix di depan API endpoint
	router := r.PathPrefix("/glodok").Subrouter()

	router.HandleFunc("", defaultHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")

	sub := router.PathPrefix("/v1").Subrouter()

	// Routes
	glodok := sub.PathPrefix("/data").Subrouter()
	glodok.HandleFunc("", s.Glodok.GetGlodok).Methods("GET")
	glodok.HandleFunc("", s.Glodok.InsertGlodok).Methods("POST")
	glodok.HandleFunc("", s.Glodok.UpdateGlodok).Methods("PUT")
	glodok.HandleFunc("", s.Glodok.DeleteGlodok).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}

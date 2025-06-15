package server

import (
	"fmt"
	"net/http"

	"github.com/hse-telescope/emailer/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	server http.Server
}

func New(conf config.Config) *Server {
	s := new(Server)
	s.server.Addr = fmt.Sprintf(":%d", conf.Port)
	s.server.Handler = s.setRouter()

	return s
}

func (s *Server) setRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	return mux
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

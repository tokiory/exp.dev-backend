package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
	addr string
	log *slog.Logger
	mux  *http.ServeMux
}

type ServerOptions struct {
	Addr string
}

func NewServer(l *slog.Logger, options ServerOptions) *Server {
	mux := http.NewServeMux()

	return &Server{
		mux:  mux,
		log: l,
		addr: options.Addr,
	}
}

func (s *Server) Apply(name string, mw func(http.Handler) http.Handler) {
	s.log.Info("Adding middleware", slog.String("middleware", name))
	originalMux := s.mux
	s.mux = http.NewServeMux()
	s.mux.Handle("/", mw(originalMux))
}

func (s *Server) Handle(pattern string, h http.HandlerFunc) {
	s.log.Info("Registered new handler", slog.String("pattern", pattern))
	s.mux.HandleFunc(pattern, h)
}

func (s *Server) Start() error {
	s.log.Info("Starting the server", slog.String("addr", s.addr))
	return http.ListenAndServe(s.addr, s.mux)
}

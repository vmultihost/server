package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	config *Config
	log    *logrus.Logger
	router *mux.Router
}

func New(config *Config) *HttpServer {
	return &HttpServer{
		config: config,
		log:    logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *HttpServer) HandleGet(path string, handler func(*logrus.Logger) http.HandlerFunc) {
	s.router.Handle(path, handler(s.log)).Methods("GET")
}

func (s *HttpServer) HandlePost(path string, handler http.HandlerFunc) {
	s.router.Handle(path, handler).Methods("POSt")
}

func (s *HttpServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	fullAddress := s.config.FullAddress()
	s.log.Infof("starting http server %s", fullAddress)

	return http.ListenAndServe(fullAddress, s.router)
}

func (s *HttpServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.logLevel)
	if err != nil {
		return err
	}

	s.log.SetLevel(level)

	return nil
}

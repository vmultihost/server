package httpserver

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ExampleServer struct {
	HttpServer
}

// todo: return interface ?
func NewExample(config *Config) *ExampleServer {
	server := *New(config)
	return &ExampleServer{
		server,
	}
}

func (s *ExampleServer) Start() error {
	s.сonfigureRouter()
	return s.HttpServer.Start()
}

func (s *ExampleServer) сonfigureRouter() {
	s.HandleGet("/hello", s.handleHello)
}

func (s *ExampleServer) handleHello(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

package cloudinit

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vmultihost/server/internal/httpserver"
)

type Server struct {
	httpserver.HttpServer
	cloudInit *CloudInit
}

func NewServer(config *httpserver.Config, cloudInit *CloudInit) *Server {
	server := *httpserver.New(config)
	return &Server{
		HttpServer: server,
		cloudInit:  cloudInit,
	}
}

func (s *Server) Start() error {
	s.HandleGet("/{instanceId}/user-data", s.userDataHandler)
	s.HandleGet("/{instanceId}/meta-data", s.metaDataHandler)

	return s.HttpServer.Start()
}

func (s *Server) metaDataHandler(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		instanceId := vars["instanceId"]

		data, err := s.cloudInit.GetMetaDataYaml(instanceId)
		if err != nil {
			log.Error("failed to get meta-data")
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		io.WriteString(w, data)
	}
}

func (s *Server) userDataHandler(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		instanceId := vars["instanceId"]

		data, err := s.cloudInit.GetUserDataYaml(instanceId)
		if err != nil {
			log.Error("failed to get user-data")
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		io.WriteString(w, data)
	}
}

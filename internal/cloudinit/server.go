package cloudinit

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	host   string
	port   uint64
	router *mux.Router
	store  map[string]*CloudInit
	log    *logrus.Logger
}

func NewServer(host string, port uint64, log *logrus.Logger) *Server {
	return &Server{
		host:   host,
		port:   port,
		router: mux.NewRouter(),
		store:  map[string]*CloudInit{},
		log:    log,
	}
}

func (s *Server) DataSource() string {
	host := fmt.Sprintf("%s:%d", s.host, s.port)
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/__dmi.chassis-serial-number__/",
	}

	return fmt.Sprintf("ds=nocloud;s=%s", u.String())
}

func (s *Server) AddCloudInit(cloudInit *CloudInit) {
	s.store[cloudInit.instanceId] = cloudInit
}

// todo: use gin
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.host, s.port)
	s.router.Handle("/{instanceId}/user-data", s.userDataHandler()).Methods("GET")
	s.router.Handle("/{instanceId}/meta-data", s.metaDataHandler()).Methods("GET")

	return http.ListenAndServe(address, s.router)
}

func (s *Server) userDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		instanceId := vars["instanceId"]
		cloudInit, ok := s.store[instanceId]
		if !ok {
			// todo: json error
			http.Error(w, "user-data not found", http.StatusNotFound)
			return
		}

		data, err := cloudInit.GetUserDataYaml()
		if err != nil {
			s.log.Error(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		io.WriteString(w, data)
	}
}

func (s *Server) metaDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		instanceId := vars["instanceId"]
		cloudInit, ok := s.store[instanceId]
		if !ok {
			// todo: json error
			http.Error(w, "meta-data not found", http.StatusNotFound)
			return
		}

		data, err := cloudInit.GetMetaDataYaml()
		if err != nil {
			s.log.Error(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		io.WriteString(w, data)
	}
}

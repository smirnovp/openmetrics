package apiserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// IConverter ...
type IConverter interface {
	GetMetrics() (string, error)
}

// APIServer ...
type APIServer struct {
	config    *Config
	logger    *logrus.Logger
	mux       *http.ServeMux
	converter IConverter
}

// New ...
func New(l *logrus.Logger, c *Config, conv IConverter) *APIServer {
	return &APIServer{
		config:    c,
		logger:    l,
		mux:       http.DefaultServeMux,
		converter: conv,
	}
}

// Run ...
func (s *APIServer) Run() error {
	server := &http.Server{
		Addr:         s.config.Addr,
		Handler:      s.mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.Routes()

	s.logger.Infof("Starting API server on %s", s.config.Addr)
	return server.ListenAndServe()
}

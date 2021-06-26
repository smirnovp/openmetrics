package apiserver

import (
	"context"
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
	httpServer *http.Server
	config     *Config
	logger     *logrus.Logger
	mux        *http.ServeMux
	converter  IConverter
	Running    chan struct{}
}

// New ...
func New(l *logrus.Logger, c *Config, conv IConverter) *APIServer {
	return &APIServer{
		config:    c,
		logger:    l,
		converter: conv,
		mux:       http.NewServeMux(),
		Running:   make(chan struct{}),
	}
}

// ConfigureAPIServer ...
func (s *APIServer) ConfigureAPIServer() error {
	if err := s.config.GetFromFile(); err != nil {
		return err
	}

	s.httpServer = &http.Server{
		Addr:         s.config.Addr,
		Handler:      s.mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.routes()

	return nil
}

// Run ...
func (s *APIServer) Run() error {

	if err := s.ConfigureAPIServer(); err != nil {
		return err
	}

	s.logger.Infof("Starting API server on %s", s.config.Addr)
	close(s.Running)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop ...
func (s *APIServer) Stop() error {
	return s.httpServer.Shutdown(context.Background())
}

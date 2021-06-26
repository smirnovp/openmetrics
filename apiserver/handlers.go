package apiserver

import (
	"net/http"
)

// Routes ...
func (s *APIServer) routes() {
	s.mux.Handle("/metrics", s.MetricsHandler())
}

// MetricsHandler ...
func (s *APIServer) MetricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metrics, err := s.converter.GetMetrics()
		if err != nil {
			s.logger.Error("Ошибка конвертирования: ", err)
			http.Error(w, "Ошибка конвертирования данных из файла", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write([]byte(metrics)); err != nil {
			s.logger.Error(err)
		}
	}
}

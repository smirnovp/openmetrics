package apiserver

import (
	"net/http"
)

// Routes ...
func (s *APIServer) Routes() {
	s.mux.Handle("/metrics", s.MetricsHandler())
}

// MetricsHandler ...
func (s *APIServer) MetricsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metrics, err := s.converter.GetMetrics()
		if err != nil {
			s.logger.Error("Ошибка конвертирования: ", err)
			http.Error(w, "Ошибка конвертироания данных из файла", http.StatusInternalServerError)
		}
		w.Write([]byte(metrics))
	}
}

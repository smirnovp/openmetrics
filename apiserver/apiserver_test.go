package apiserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type MockConverter struct{}

func (m *MockConverter) GetMetrics() (string, error) {
	return "currencies{name:usd} 70\ncurrencies{name:eur} 80\n", nil
}

type MockConverterError struct{}

func (m *MockConverterError) GetMetrics() (string, error) {
	return "", fmt.Errorf("Error")
}

func TestAPIServer(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/metrics", nil)
	if err != nil {
		t.Error(err)
	}

	server := New(logrus.New(), NewConfig("config/config.toml"), &MockConverter{})

	err = server.ConfigureAPIServer()
	assert.NotEqual(t, nil, err, "Ошибка должна быть nil")

	server.mux.ServeHTTP(w, r)
	assert.Equal(t, "currencies{name:usd} 70\ncurrencies{name:eur} 80\n", w.Body.String(), "Должны быть равны")

	server.converter = &MockConverterError{}
	w = httptest.NewRecorder()
	server.mux.ServeHTTP(w, r)
	assert.Equal(t, http.StatusInternalServerError, w.Code, fmt.Sprintf("Код должен быть: %d", http.StatusInternalServerError))

}

func TestGetDataFromFile(t *testing.T) {
	c := NewConfig("testdata/config.toml")
	err := c.GetFromFile()
	assert.Equal(t, nil, err, "err должна быть nil")
	assert.Equal(t, ":8080", c.Addr, "Должны быть равны")

	c.FileName = "testdata/badconfig.toml"
	err = c.GetFromFile()
	assert.NotEqual(t, nil, err, "err не должна быть nil")
}

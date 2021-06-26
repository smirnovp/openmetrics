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

func TestListenAndServeFail(t *testing.T) {

	server := New(logrus.New(), NewConfig("testdata/configBadAddr.toml"), &MockConverter{})
	err := server.Run()
	assert.NotEqual(t, nil, err, "Должна быть ошибка с занятым портом")
}

func TestAPIServerRUN(t *testing.T) {

	server := New(logrus.New(), NewConfig("../config/config.toml"), &MockConverter{})

	serverStopped := make(chan struct{})
	errorc := make(chan struct{})

	go func() {
		err := server.Run()
		if err != nil {
			t.Error("Ошибка запуска сервера: ", err)
			errorc <- struct{}{}
		}
		defer close(serverStopped)
	}()

	select {
	case <-server.Running: // ждем пока запустится горутина запускающая сервер:
	case <-errorc: // если сервер не запустился, то выходим
		return
	}

	err := server.Stop()
	if err != nil {
		t.Error(err)
	}

	<-serverStopped // ждем пока сервер остановится и Go-рутина закроется
}

func TestAPIServerRUNFail(t *testing.T) {

	server := New(logrus.New(), NewConfig("../config/confi.toml"), &MockConverter{})

	serverStopped := make(chan struct{})
	errorc := make(chan struct{})

	go func() {
		err := server.Run()
		if err != nil {
			assert.NotEqual(t, nil, err, "Должна быть ошибка запуска сервера")
			errorc <- struct{}{}
		}
		defer close(serverStopped)
	}()

	select {
	case <-server.Running: // ждем пока запустится горутина запускающая сервер:
	case <-errorc: // если сервер не запустился, то выходим
		return
	}

	err := server.Stop()
	if err != nil {
		t.Error(err)
	}

	<-serverStopped // ждем пока сервер остановится и Go-рутина закроется
}

func TestAPIServer(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/metrics", nil)
	if err != nil {
		t.Error(err)
	}

	server := New(logrus.New(), NewConfig("../config/config.toml"), &MockConverter{})

	err = server.ConfigureAPIServer()
	if err != nil {
		t.Error("Ошибка конфигурации сервера: ", err)
	}

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
	assert.Equal(t, ":8082", c.Addr, "Должны быть равны")

	c.FileName = "testdata/badconfig.toml"
	err = c.GetFromFile()
	assert.NotEqual(t, nil, err, "err не должна быть nil")
}

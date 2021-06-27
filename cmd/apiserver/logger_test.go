package main

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConfigureLogger(t *testing.T) {
	l := logrus.New()
	err := configureLogger(l, "testdata/config.toml")
	assert.Equal(t, nil, err, "Ошибки быть не должно")
	assert.Equal(t, logrus.DebugLevel, l.GetLevel(), "loglevel должен быть `debug`")

	err = configureLogger(l, "testdata/badconfig.toml")
	assert.NotEqual(t, nil, err, "Должна быть ошибка считывания конфигурации")

	err = configureLogger(l, "testdata/badlevelconfig.toml")
	assert.NotEqual(t, nil, err, "Должна быть ошибка парсинга loglevel")
}

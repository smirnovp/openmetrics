package main

import (
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func configureLogger(l *logrus.Logger, f string) error {

	viper.SetConfigFile(path.Clean(f))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.SetDefault("logger.level", "debug")
	ls := viper.GetString("logger.level")

	level, err := logrus.ParseLevel(ls)
	if err != nil {
		return err
	}

	l.SetLevel(level)

	return nil
}

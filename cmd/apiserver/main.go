package main

import (
	"flag"
	"log"
	"openmetrics/apiserver"
	"openmetrics/converter"

	"github.com/sirupsen/logrus"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "config/config.toml", "Server configuration file")
	flag.Parse()

	logger := logrus.New()
	if err := configureLogger(logger, configFile); err != nil {
		log.Fatal("Не смог сконфигурировать logger: ", err)
	}

	converter := converter.NewFileConverter("currencies.yaml")

	config := apiserver.NewConfig(configFile)
	apiserver := apiserver.New(logger, config, converter)

	err := apiserver.Run()
	if err != nil {
		log.Fatal(err)
	}
}

package apiserver

import (
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	FileName string
	Addr     string
}

// NewConfig ...
func NewConfig(f string) *Config {
	return &Config{
		FileName: f,
		Addr:     ":8080",
	}
}

// GetFromFile ...
func (c *Config) GetFromFile() error {
	viper.SetConfigFile(c.FileName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("Addr", ":8080")
	c.Addr = viper.GetString("apiserver.addr")

	return nil
}

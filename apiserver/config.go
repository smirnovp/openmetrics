package apiserver

import (
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Addr string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Addr: ":8080",
	}
}

// GetFromFile ...
func (c *Config) GetFromFile(f string) error {
	viper.SetConfigFile(f)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("Addr", ":8080")
	c.Addr = viper.GetString("Addr")

	return nil
}

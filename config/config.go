package config

import (
	"github.com/gojek-engineering/goconfig"
	"log"
)

type Config struct {
	goconfig.BaseConfig
}

var config *Config

func init() {
	log.Print("Loading the config!")

	config = &Config{}
	config.Load()
}

func Port() int {
	return config.GetIntValue("port")
}

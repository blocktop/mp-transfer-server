package config

import (
	"github.com/blocktop/mp-common/config"
	"github.com/caarlos0/env"
)

// Config holds system configuration values.
type Config struct {
	config.BaseConfig
	TransferServer         string `env:"MP_TRANSFER_SERVER"`
	TransferServerInfoPath string `env:"MP_TRANSFER_SERVER_INFO_PATH"`
}

var cfg *Config

func init() {
	cfg = &Config{}
	cfg.Parse()
}

func GetConfig() *Config {
	return cfg
}

func (c *Config) Parse() {
	c.BaseConfig.Parse()

	err := env.Parse(c)
	if err != nil {
		panic(err)
	}
}

package config

import (
	"github.com/coscms/webcore/library/config"
	"github.com/coscms/webcore/library/config/extend"
)

func init() {
	extend.Register(`download`, func() interface{} {
		return &Config{}
	})
}

func Get() *Config {
	cfg, _ := config.MustGetConfig().Extend.Get(`download`).(*Config)
	if cfg == nil {
		cfg = &Config{}
	}
	return cfg
}

type Config struct {
	SavePath string `json:"savePath"`
}

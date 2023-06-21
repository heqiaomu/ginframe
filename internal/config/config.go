package config

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Port  int  `json:"port"`
	Pprof bool `json:"pprof"`
}

var cfg *Config

func InitConfig(filePath string) (*Config, error) {
	cfg = &Config{}
	if err := conf.LoadConfig(filePath, cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to load config")
	}
	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

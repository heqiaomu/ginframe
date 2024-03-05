package config

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Port   int    `json:"port" yaml:"port"`
	Pprof  bool   `json:"pprof" yaml:"pprof"`
	Prefix string `json:"prefix" yaml:"prefix"`
}

var cfg *Config

func InitConfig(filePath string) (*Config, error) {
	cfg = &Config{}
	if err := conf.Load(filePath, cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to load config")
	}
	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

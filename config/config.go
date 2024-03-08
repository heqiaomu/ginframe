package config

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Server Server `json:"server" yaml:"server"`
	Rbac   Rbac   `json:"rbac" yaml:"rbac"`
}

type Rbac struct {
	Enabled    bool   `json:"enabled" yaml:"enabled"`
	ModelPath  string `json:"model_path" yaml:"model_path"`
	PolicePath string `json:"police_path" yaml:"police_path"`
	Redis      DB     `json:"redis" yaml:"redis"`
}

type DB struct {
	DBType   string `json:"db_type" yaml:"db_type"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
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

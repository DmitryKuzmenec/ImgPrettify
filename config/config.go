package config

import (
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

// Config - struct of config
type Config struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	} `envconfig:"MYSQL"`
	Log struct {
		Level string
	} `envconfig:"LOG"`
}

// Init - get config
func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	var conf Config
	if err := envconfig.Init(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

package config

import (
	"github.com/BurntSushi/toml"
	"pyke/logger"
)

const DefaultConfigPath = "./pyke_config.toml"

type Config struct {
	Host string
	Port int
	Mode string
}

// --------- constructor ---------

func New(path string) (*Config, error) {
	c := Config{}
	_, err := toml.DecodeFile(path, &c)
	if nil != err {
		logger.PykeError.Println(err)
		return nil, err
	}
	return &c, nil
}

// --------- instance ---------

var PykeConfigInstant *Config

func init() {
	logger.PykeInfo.Println("PykeConfigInstant init")
	PykeConfigInstant, _ = New(DefaultConfigPath)
}

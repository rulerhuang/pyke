package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

const CONFIF_PATH = "./pyke_config.toml"

type Config struct {
	Host string
	Port int
	Mode string
}

func getConfig(path string) (*Config, error) {
	c := Config{}
	_, err := toml.DecodeFile(path, &c)
	if nil != err {
		log.Printf("error:%s", err.Error())
		return nil, err
	}
	log.Printf("config:%+v.\n", c)
	return &c, nil
}

// --------- constructor ---------

func New() (*Config, error) {
	return getConfig(CONFIF_PATH)
}

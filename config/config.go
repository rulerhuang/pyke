package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
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
		log.Printf("error:%s", err.Error())
		return nil, err
	}
	return &c, nil
}

// --------- instance ---------

var ConfigInstant *Config

func init() {
	fmt.Println("ConfigInstant init")
	ConfigInstant, _ = New(DefaultConfigPath)
}

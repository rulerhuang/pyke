package utils

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Host  string
	Port  int
	Debug bool
}

func GetConfig(path string) *Config {
	c := Config{}
	_, err := toml.DecodeFile(path, &c)
	if nil != err {
		log.Printf("error:%s", err.Error())
	}
	log.Printf("config:%+v.\n", c)
	return &c
}

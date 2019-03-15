package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"pyke/rule"
	"sync"
)

type JsonFileStorage struct {
	rules []rule.Rule
	mutex sync.RWMutex
	count int
}

// --------- methods ---------

func (c *JsonFileStorage) Set(r *rule.Rule) error {
	return nil
}

func (c *JsonFileStorage) Load() error {
	return nil
}

func (c *JsonFileStorage) Save() error {
	return nil
}

func Load() error {
	f, err := ioutil.ReadFile("./debug_rules.json")
	if nil != err {
		log.Printf("error:%s", err.Error())
		return err
	}

	c := make([]rule.Rule, 10)
	err = json.Unmarshal(f, &c)
	if nil != err {
		log.Printf("error:%s", err.Error())
		return err
	}

	fmt.Printf("%+v\n", c)
	return nil
}

func Save(c *rule.Rule) error {
	jsonByteData, err := json.Marshal(*c)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	err = ioutil.WriteFile("./new_debug_rules.json", jsonByteData, 0644)
	if nil != err {
		log.Printf("error:%s", err.Error())
		return err
	}
	return nil
}

// --------- constructor ---------

func newJsonFileStorage() *JsonFileStorage {
	return &JsonFileStorage{make([]rule.Rule, 0), sync.RWMutex{}, 0}
}

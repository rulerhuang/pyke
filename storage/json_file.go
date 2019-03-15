package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"pyke/rule"
	"sync"
)

const DEFAULT_JSON_FILE_PATH = "./debug_rules.json"

type JsonFileStorage struct {
	Rules []rule.Rule
	Mutex sync.RWMutex
	Path  string
	Count int
}

// --------- methods ---------

func (c *JsonFileStorage) Load() error {
	f, err := ioutil.ReadFile(c.Path)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	r := make([]rule.Rule, 0)
	err = json.Unmarshal(f, &r)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return err
	}

	c.Rules = r
	c.Count = len(c.Rules)
	return nil
}

func (c *JsonFileStorage) Set(r *rule.Rule) error {
	c.Rules = append(c.Rules, *r)
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

func newJsonFileStorage(fp string) *JsonFileStorage {
	if fp == "" {
		fp = DEFAULT_JSON_FILE_PATH
	}

	return &JsonFileStorage{
		Rules: make([]rule.Rule, 0),
		Mutex: sync.RWMutex{},
		Path:  fp,
		Count: 0}
}

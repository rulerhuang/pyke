package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"pyke/rule"
	"sync"
)

const DefaultJsonFilePath = "./demo_rules.json"

type JsonFileStorage struct {
	Rules []rule.Rule
	Mutex sync.RWMutex
	Path  string
	Count int
}

// --------- methods ---------

func (c *JsonFileStorage) Load() (int, error) {
	f, err := ioutil.ReadFile(c.Path)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return 0, err
	}

	r := make([]rule.Rule, 0)
	err = json.Unmarshal(f, &r)
	if err != nil {
		log.Printf("error:%s", err.Error())
		return 0, err
	}

	c.Rules = r
	c.Count = len(c.Rules)
	return len(c.Rules), nil
}

func (c *JsonFileStorage) Save() error {
	return nil
}

func (c *JsonFileStorage) Set(r *rule.Rule) error {
	c.Rules = append(c.Rules, *r)
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
		fp = DefaultJsonFilePath
	}

	return &JsonFileStorage{
		Rules: make([]rule.Rule, 0),
		Mutex: sync.RWMutex{},
		Path:  fp,
		Count: 0}
}

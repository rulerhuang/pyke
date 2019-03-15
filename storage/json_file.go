package storage

import (
	"encoding/json"
	"io/ioutil"
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
		return 0, err
	}

	r := make([]rule.Rule, 0)
	err = json.Unmarshal(f, &r)
	if err != nil {
		return 0, err
	}

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Rules = r
	c.Count = len(c.Rules)
	return len(c.Rules), nil
}

func (c *JsonFileStorage) Get() ([]rule.Rule, error) {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()

	return c.Rules, nil
}

func (c *JsonFileStorage) Set(r *rule.Rule) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Rules = append(c.Rules, *r)
	c.Count += 1
	return nil
}

func (c *JsonFileStorage) Save() error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	jsonByteData, err := json.Marshal(c.Rules)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.Path, jsonByteData, 0644)
	if nil != err {
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

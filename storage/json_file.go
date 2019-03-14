package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"pyke/rule"
)

func Load() {
	f, err := ioutil.ReadFile("./debug_rules.json")
	if nil != err {
		log.Printf("error:%s", err.Error())
	}

	c := make([]rule.Rule, 10)
	err = json.Unmarshal(f, &c)
	if nil != err {
		log.Printf("error:%s", err.Error())
	}

	fmt.Printf("%+v\n", c)
}

func Save(c *rule.Rule) {
	jsonByteData, err := json.Marshal(*c)
	if err != nil {
		log.Printf("error:%s", err.Error())
	}

	err = ioutil.WriteFile("./new_debug_rules.json", jsonByteData, 0644)
	if nil != err {
		log.Printf("error:%s", err.Error())
	}
}

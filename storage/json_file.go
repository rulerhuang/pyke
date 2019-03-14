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

func dump() {

}

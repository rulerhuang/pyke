package rule

import "sync"

type MemeRuleMap struct {
	c     map[string][]Rule
	mutex sync.RWMutex
}

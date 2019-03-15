package rule

import "sync"

type Topic struct {
	key   string
	rules map[string][]Rule
	mutex sync.RWMutex
}

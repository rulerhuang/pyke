package rule

import "sync"

type Topic struct {
	rules map[string][]Rule
	mutex sync.RWMutex
}

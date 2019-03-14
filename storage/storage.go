package storage

import "pyke/rule"

type Storage interface {
	Load() error
	Save(rule *rule.Rule) error
	dump() error
}

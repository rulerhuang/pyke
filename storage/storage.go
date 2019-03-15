package storage

import "pyke/rule"

const (
	JSON_MODE   = "json_file"
	BOLTDB_MODE = "bolt_db"
)

type Storage interface {
	Load() error
	Set(rule *rule.Rule) error
	Save() error
}

func New(tp string) Storage {
	var s Storage
	if JSON_MODE == tp {
		s = newJsonFileStorage(tp)
	} else if BOLTDB_MODE == tp {
		s = nil
	}

	if s == nil {
		panic("load rules failed")
	}

	return s
}

package storage

import "pyke/rule"

const (
	JsonMode   = "json_file"
	BoltDBMode = "bolt_db"
)

type Storage interface {
	Load() (int, error)
	Save() error
	Get() ([]rule.Rule, error)
	Set(rule *rule.Rule) error
}

// --------- constructor ---------

func New(tp string, pt string) Storage {
	var s Storage
	if JsonMode == tp {
		s = newJsonFileStorage(pt)
	} else if BoltDBMode == tp {
		s = nil
	}

	if s == nil {
		panic("load rules failed")
	}

	return s
}

// --------- instance ---------

var RuleStorage = New(JsonMode, DefaultJsonFilePath)

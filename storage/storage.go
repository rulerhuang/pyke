package storage

import (
	"fmt"
	"pyke/rule"
)

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

func New(mode string, filePath string) Storage {
	var s Storage
	if JsonMode == mode {
		s = newJsonFileStorage(filePath)
	} else if BoltDBMode == mode {
		s = nil
	}

	if s == nil {
		panic("load rules failed")
	}

	return s
}

// --------- instance ---------

var RuleStorageInstant Storage

func init() {
	fmt.Println("RuleStorageInstant init")
	RuleStorageInstant = New(JsonMode, DefaultJsonFilePath)
	// get init rules
	RuleStorageInstant.Load()
}

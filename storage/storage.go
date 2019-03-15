package storage

import (
	"pyke/logger"
	"pyke/rule"
)

const (
	JsonMode   = "json_file"
	BoltDBMode = "bolt_db"
)

type Storage interface {
	Load() (int, error)
	Get() ([]rule.Rule, error)
	Set(rule *rule.Rule) error
	Save() error
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

var PykeStorageInstant Storage

func init() {
	logger.Info.Println("PykeStorageInstant init")
	PykeStorageInstant = New(JsonMode, DefaultJsonFilePath)
	// get init rules
	_, _ = PykeStorageInstant.Load()
}

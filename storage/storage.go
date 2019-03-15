package storage

import "pyke/rule"

type Storage interface {
	Load() error
	Set(rule *rule.Rule) error
	Save() error
}

func New(tp string) Storage {
	var s Storage
	if "json_file" == tp {
		s = newJsonFileStorage()
	}

	if nil == s {
		panic("load rules failed")
	}

	return s
}

package storage

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	j := newJsonFileStorage("../debug_rules.json")
	err := j.Load()
	if err != nil {
		t.Fatal("load error")
	}
	fmt.Printf("%v:%d", j.Rules, j.Count)
}

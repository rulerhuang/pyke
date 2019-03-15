package storage

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	j := newJsonFileStorage("../demo_rules.json")
	num, err := j.Load()
	if err != nil {
		t.Fatal("load error")
	}
	fmt.Printf("%v:%d", j.Rules, num)
}

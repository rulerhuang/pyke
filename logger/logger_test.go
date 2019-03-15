package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Info.Println("TestInfo")
}

func TestError(t *testing.T) {
	Error.Println("TestError")
}

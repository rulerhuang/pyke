package logger

import (
	"testing"
)

func TestInfo(t *testing.T) {
	PykeInfo.Println("TestInfo")
}

func TestError(t *testing.T) {
	PykeError.Println("TestError")
}

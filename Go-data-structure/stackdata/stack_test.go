package stackdata_test

import (
	"testing"

	"GoAlgo/stackdata"
)

func TestNewStack(t *testing.T) {
	stack := stackdata.NewStack[int]()
	if stack == nil {
		t.Error("NewStack() returned nil")
	}
}

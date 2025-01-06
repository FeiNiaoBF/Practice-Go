package foorpg

import (
	"bytes"
	"testing"
	"time"
)

func TestDone(t *testing.T) {
	c := &Call{
		Done: make(chan *Call, 1),
	}
	go func(call *Call) {
		time.Sleep(time.Second)
		call.done()
		t.Log("call done")
	}(c)

	select {
	case <-c.Done:
		t.Log("done")
	case <-time.After(time.Second * 2):
		t.Error("timeout")
	}
}

// Test Client

func TestRegisterCall(t *testing.T) {
	t.Log("register")
	buf := &bufferWrapper{Buffer: new(bytes.Buffer)}
	client, err := NewClient(buf, DefaultOption)
	if err != nil {
		t.Errorf("new client failed: %v", err)
	}

	call := &Call{
		Seq:           1,
		ServiceMethod: "foo/bar",
		Args:          "hello",
		Reply:         "world",
	}
	seq, err := client.registerCall(call)
	t.Log("done")
	if err != nil {
		t.Error(err)
	}
	if seq != call.Seq {
		t.Errorf("expect %d but got %d", call.Seq, seq)
	}
}

package codec

import (
	"net"
	"reflect"
	"testing"
)

// test NewGobCodec ReadHeader
func TestGobCodec_ReadHeader(t *testing.T) {
	// create a pipe
	pair, _ := net.Pipe()
	// create a new GobCodec
	codec := NewGobCodec(pair)
	// create a mock header
	h := &Header{
		ServiceMethod: "Service.Method",
		Seq:           1,
	}
	// encode the header
	_ = codec.Write(h, nil)
	// create a new header
	h1 := &Header{}
	// read the header
	err := codec.ReadHeader(h1)
	if err != nil {
		t.Errorf("ReadHeader failed: %v", err)
	}
	if !reflect.DeepEqual(h, h1) {
		t.Errorf("ReadHeader failed: %v != %v", h, h1)
	}
}

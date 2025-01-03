package codec

import (
	"bytes"
	"reflect"
	"testing"
)

type bufferWrapper struct {
	*bytes.Buffer
}

func (b *bufferWrapper) Close() error {
	return nil
}

// test NewGobCodec ReadHeader
func TestGobCodec_ReadHeaderAndWrite(t *testing.T) {
	// create a buf
	buf := &bufferWrapper{Buffer: new(bytes.Buffer)}
	// create a new GobCodec
	codec := NewGobCodec(buf)
	// create a mock header
	h := &Header{
		ServiceMethod: "Service.Method",
		Seq:           1,
		Error:         "",
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

// test NewGobCodec ReadBody

// test NewGobCodec Write
func TestGobCodec_Write(t *testing.T) {
	// create a buf
	buf := &bufferWrapper{Buffer: new(bytes.Buffer)}
	// create a new GobCodec
	codec := NewGobCodec(buf)
	// create a mock header
	h := &Header{
		ServiceMethod: "Service.Method",
		Seq:           1,
		Error:         "",
	}
	t.Run("Write header", func(t *testing.T) {
		// write the header
		err := codec.Write(h, nil)
		if err != nil {
			t.Errorf("Write failed: %v", err)
		}
	})
	t.Run("Write body", func(t *testing.T) {
		// write the body
		// create a mock body
		type TestBody struct {
			Name string
			Age  int
		}
		body := &TestBody{
			Name: "test",
			Age:  18,
		}
		err := codec.Write(h, body)
		if err != nil {
			t.Errorf("Write failed: %v", err)
		}
	})
}

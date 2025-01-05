package foorpg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fooprc/codec"
	"log"
	"net"
	"testing"
	"time"
)

type bufferWrapper struct {
	*bytes.Buffer
}

func (b *bufferWrapper) Close() error {
	return nil
}

func TestReadRequest(t *testing.T) {
	h := &codec.Header{
		ServiceMethod: "foo/bar",
		Seq:           1,
	}

	t.Run("read head", func(t *testing.T) {
		r := &bufferWrapper{new(bytes.Buffer)}

		s := &Server{}
		// create a mock codec
		mockCodec := codec.NewGobCodec(r)
		_ = mockCodec.Write(h, nil)
		h1, err := s.readRequestHeader(mockCodec)
		if err != nil {
			t.Errorf("readRequestHeader failed: %v", err)
		}
		if h1.ServiceMethod != h.ServiceMethod {
			t.Errorf("readRequestHeader failed: %v != %v", h1.ServiceMethod, h.ServiceMethod)
		}
	})

	t.Run("read request", func(t *testing.T) {
		r := &bufferWrapper{new(bytes.Buffer)}

		s := &Server{}
		// create a mock codec
		mockCodec := codec.NewGobCodec(r)
		_ = mockCodec.Write(h, nil)
		req, err := s.readRequest(mockCodec)
		if err != nil {
			t.Errorf("readRequest failed: %v", err)
		}

		if req.h.ServiceMethod != h.ServiceMethod {
			t.Errorf("readRequest failed: %v != %v", req.h.ServiceMethod, h.ServiceMethod)
		}
	})
}

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	Accept(l)
}

func TestServerCanHandleRequest(t *testing.T) {
	addr := make(chan string)
	go startServer(addr)
	// 模拟一个客户端连接到rpc server
	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)
	// send options
	// 发送默认的Option
	_ = json.NewEncoder(conn).Encode(DefaultOption)
	// 创建一个新的GobCodec（打包）
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 5; i++ {
		// mock header
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("foorpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}

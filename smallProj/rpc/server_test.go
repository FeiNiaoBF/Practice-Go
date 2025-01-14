package foorpg

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"fooprc/codec"
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
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)
	// 使用 client
	client, _ := Dial("tcp", <-addr)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	// send request & receive response
	for i := 0; i < 5; i++ {
		// mock header
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("foo %d", i)
			var reply string
			if err := client.Call("foo", args, &reply); err != nil {
				log.Fatalf("call foo failed: %v", err)
			}
			log.Println("reply:", reply)
		}(i)
	}

	wg.Wait()
}

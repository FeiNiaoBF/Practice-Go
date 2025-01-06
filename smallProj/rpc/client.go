package foorpg

import (
	"errors"
	"io"
	"sync"

	"fooprc/codec"
)

// Call represents an active RPC.
type Call struct {
	Seq           uint64
	ServiceMethod string      // format "<service>.<method>"
	Args          interface{} // arguments to the function
	Reply         interface{} // reply from the function
	Error         error       // if error occurs, it will be set
	Done          chan *Call  // for rpc server to send back the call
}

// done() sends call to done channel
func (c *Call) done() {
	c.Done <- c
}

// Client represents an RPC Client.
// There may be multiple outstanding Calls associated
// with a single Client, and a Client may be used by
// multiple goroutines simultaneously.
type Client struct {
	cc       codec.Codec // codec to encode/decode requests
	opt      *Option
	sending  sync.Mutex   // protect following
	header   codec.Header // request header
	mu       sync.Mutex   // protect following
	seq      uint64       // sequence number of request
	pending  map[uint64]*Call
	closing  bool // user has called Close
	shutdown bool // server has told us to stop
}

var _ io.Closer = (*Client)(nil)

// ErrShutdown is returned when try to call a method on a client that is shut down
var ErrShutdown = errors.New("connection is shut down")

// Close the connection to the RPC server.
func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

// IsAvailable returns true if the client is available
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown
}

// NewClient returns a new Client.
func NewClient(conn io.ReadWriteCloser, opt *Option) (*Client, error) {
	return nil, nil
}

// registerCall adds call to pending list
func (client *Client) registerCall(call *Call) (uint64, error) {
	return 0, nil
}
// removeCall deletes call from pending list by seq

// terminateCalls cleans up calls in pending list

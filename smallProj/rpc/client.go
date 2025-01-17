package foorpg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
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
	// check codec type
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error: ", err)
		return nil, err
	}
	// send options to server
	// will be opt send to server by json
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	// client codec is created by the codec function
	// will be used to encode/decode requests
	// and receive responses from the server
	return newClientCodec(f(conn), opt), nil
}

func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1, // seq starts with 1, 0 means invalid call
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client
}

// receive reads from the connection and handles the response
// call 不存在，可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了。
// call 存在，但服务端处理出错，即 h.Error 不为空。
// call 存在，服务端处理正常，那么需要从 body 中读取 Reply 的值。
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}
		// get call by seq
		call := client.removeCall(h.Seq)
		switch {
		case call == nil:
			// call does not exist
			err = client.cc.ReadBody(nil)
		case h.Error != "":
			// call exists, but server returns error
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done()
		default:
			// call exists, server returns normal response
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	// error occurs, clean up calls
	client.terminateCalls(err)
}

// registerCall adds call to pending list
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}
	// sequential number of the call request is the key
	call.Seq = client.seq
	client.pending[call.Seq] = call
	// not unique, increment
	client.seq++

	return call.Seq, nil
}

// removeCall deletes call from pending list by seq
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// terminateCalls cleans up calls in pending list
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

// parseOptions parses the options
func parseOptions(opts ...*Option) (*Option, error) {
	// if opts is nil or pass nil as parameter
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// Dial connects to an RPC server at the specified network address.
func Dial(network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		log.Printf("rpc client: dial option error: %v", err)
		return nil, err
	}
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Printf("rpc client: dial error: %v", err)
		return nil, err
	}
	defer func() {
		if client == nil {
			_ = conn.Close()
		}
	}()
	return NewClient(conn, opt)
}

// send sends the request to the server
func (client *Client) send(call *Call) {
	// protect sending
	client.sending.Lock()
	defer client.sending.Unlock()
	// register call
	seq, err := client.registerCall(call)
	if err != nil {
		call.Error = err
		call.done()
		return
	}
	// prepare request header
	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""
	// encode request header
	if err = client.cc.Write(&client.header, call.Args); err != nil {
		call := client.removeCall(seq)
		if call != nil {
			call.Error = err
			call.done()
		}
	}
}

// Go invokes the function asynchronously.
// It returns the Call structure representing the invocation.
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, 10)
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	client.send(call)
	return call
}

// Call invokes the named function, waits for it to complete,
// and returns its error status.
func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}

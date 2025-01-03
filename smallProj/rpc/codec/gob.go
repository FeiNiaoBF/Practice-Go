package codec

import (
	"bufio"
	"encoding/gob"
	"io"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

var _ Codec = (*GobCodec)(nil)

func (c *GobCodec) ReadHeader(h *Header) error {
	return nil
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return nil
}

func (c *GobCodec) Write(h *Header, body interface{}) error {
	return nil
}

func (c *GobCodec) Close() error {
	return nil
}

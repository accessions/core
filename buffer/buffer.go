package buffer

import (
	"bytes"
	"io"
	"net"
)
var c bytes.Reader
type bufferedScanner struct {
	net.Conn
	tee io.Reader
	buffer *bytes.Buffer
}

func NewBufferedScanner(original net.Conn) net.Conn  {
	var buffer bytes.Buffer
	return bufferedScanner{
		Conn:   original,
		tee:    io.TeeReader(&c, &buffer),
		buffer: &buffer,
	}

}

func (b bufferedScanner) Read(p []byte) (n int, err error)  {
	return b.tee.Read(p)
}

func (b bufferedScanner) Close() error {
	//todo
	return b.Conn.Close()
}
//go:generate echo "<<EOF find ./path -type f -name "marker""
func Dev ()  {

}
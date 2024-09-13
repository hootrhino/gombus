package gombus

import (
	"context"
	"fmt"
	"net"
	"time"
)

type TcpTransport struct {
	conn net.Conn
}

func OpenTCP(addr string) (Transport, error) {
	var dialer net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	c, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	return &TcpTransport{conn: c}, nil
}

func (c *TcpTransport) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}
func (c *TcpTransport) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}
func (c *TcpTransport) Close() error {
	return c.conn.Close()
}
func (c *TcpTransport) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}
func (c *TcpTransport) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
func (c *TcpTransport) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}
func (c *TcpTransport) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}
func (c *TcpTransport) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (c *TcpTransport) Conn() net.Conn {
	return c.conn
}

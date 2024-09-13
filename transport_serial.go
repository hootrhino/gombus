package gombus

import (
	serial "github.com/hootrhino/goserial"
)

type SerialTransport struct {
	conn serial.Port
}

func OpenSerial(config serial.Config) (Transport, error) {
	transport, err := serial.Open(&config)
	if err != nil {
		return nil, err
	}
	return &SerialTransport{conn: transport}, nil
}
func (c *SerialTransport) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}
func (c *SerialTransport) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}
func (c *SerialTransport) Close() error {
	return c.conn.Close()
}

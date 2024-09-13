package gombus

import (
	"io"
)

type Transport interface {
	io.ReadWriteCloser
}

package gracenet

import (
	"net"
)

type GraceTCPConn struct {
	net.Conn
}

func (w GraceTCPConn) Close() error {
	httpWg.Done()
	return w.Conn.Close()
}

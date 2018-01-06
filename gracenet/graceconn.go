package gracenet

import (
	"net"
	"sync"
)

var (
	addCount = 0
	mCount   = 0
)

type GraceTCPConn struct {
	net.Conn
	wg *sync.WaitGroup
}

func (c GraceTCPConn) Close() error {
	mCount += 1
	// fmt.Printf("close %v pid is: %v\n", mCount, os.Getpid())
	// c.wg.Done()
	return c.Conn.Close()
}

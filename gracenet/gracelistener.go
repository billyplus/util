package gracenet

import (
	"fmt"
	"net"
	"os"
	"sync"
	"syscall"
)

type GraceListener interface {
	// Accept() (net.Conn, error)
	// Close() error
	// Addr() net.Addr
	File() (f *os.File, err error)
}

type GraceTCPListener struct {
	net.Listener
	stop    chan error
	stopped bool
	wg      *sync.WaitGroup
}

func (gl *GraceTCPListener) Accept() (c net.Conn, err error) {
	c, err = gl.Listener.Accept()
	if err != nil {
		return
	}

	c = GraceTCPConn{Conn: c, wg: gl.wg}
	addCount += 1
	fmt.Printf("add %v pid is: %v\n", addCount, os.Getpid())

	// gl.wg.Add(1)
	return
}

func NewGraceTCPListener(l net.Listener) (gl *GraceTCPListener) {
	gl = &GraceTCPListener{Listener: l, stop: make(chan error), wg: new(sync.WaitGroup)}
	go func() {
		_ = <-gl.stop
		gl.stopped = true
		gl.stop <- gl.Listener.Close()
	}()
	return
}

func (gl *GraceTCPListener) Close() error {
	if gl.stopped {
		return syscall.EINVAL
	}
	gl.stop <- nil
	return <-gl.stop
}

func (gl *GraceTCPListener) File() (*os.File, error) {
	tl := gl.Listener.(GraceListener)
	// fl, _ := tl.File()
	return tl.File()
}

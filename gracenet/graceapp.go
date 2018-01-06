package gracenet

import (
	"net"
)

type GraceServer interface {
	Listener() net.Listener
}

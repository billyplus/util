package ylua

import (
	"bufio"
	"io"
)

type Scaner struct {
	r *bufio.Reader
}

func NewScaner(r io.Reader) *Scaner {
	return &Scaner{r: bufio.NewReader(r)}
}

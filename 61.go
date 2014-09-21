package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for i := 0; i < n; i++ {
		if p[i] >= 'A' && p[i] < 'N' || p[i] >= 'a' && p[i] < 'n' {
			p[i] += 13
		} else if p[i] >= 'N' && p[i] < 'Z' || p[i] >= 'n' && p[i] < 'z' {
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

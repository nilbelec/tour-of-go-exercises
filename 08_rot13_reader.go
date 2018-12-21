package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, error := r.r.Read(b)
	if error != nil {
		return 0, error
	}
	for i := 0; i < n; i++ {
		b[i] -= 13
	}
	return n, nil
}

func exercise08() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

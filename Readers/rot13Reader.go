package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (x rot13Reader) Read(s []byte) (int, error) {
	n, err := x.r.Read(s)

	for i := 0; i < n; i++ {
		if s[i] >= 97 && s[i] <= 122 {
			s[i] += 13
			if s[i] >= 122 {
				s[i] = s[i] - 122
				s[i] = 96 + s[i]
			}
		} else if s[i] >= 65 && s[i] <= 90 {
			s[i] += 13
			if s[i] >= 90 {
				s[i] = s[i] - 90
				s[i] = 64 + s[i]
			}
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
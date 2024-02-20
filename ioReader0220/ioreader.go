package main

import (
	"fmt"
	"io"
)

type AlphaReader struct {
	str string
	cur int
}

func newReader(src string) *AlphaReader {
	return &AlphaReader{str: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *AlphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.str) {
		return 0, io.EOF
	}
	x := len(a.str) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}
	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.str[a.cur]); char != 0 {
			buf[n] = char
		}
		n++
		a.cur++
	}
	copy(p, buf)
	return n, nil
}

func main() {
	reader := newReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))

	}
	fmt.Println()
}

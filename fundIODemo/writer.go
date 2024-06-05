package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var (
		buf1 bytes.Buffer
		buf2 bytes.Buffer
	)

	// writing into mw will write to both buf1 and buf2
	mw := io.MultiWriter(&buf1, &buf2)

	// r is the source of data(Reader)
	r := strings.NewReader("some io.Reader stream to be read")

	// write to mw from r
	io.Copy(mw, r)

	fmt.Println("data inside buffer1 :", buf1.String())
	fmt.Println("data inside buffer2 :", buf2.String())
}

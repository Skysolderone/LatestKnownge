package test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Fprintf
func TestFprintf(t *testing.T) {
	file, err := os.Create("ostest.txt")
	if err != nil {
		t.Log(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "Hello,%s", "world")
}

// bufio.NewWriter
func TestNewWriter(t *testing.T) {
	file, err := os.Create("ostestNewwriter.txt")
	if err != nil {
		t.Log(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("hello newreadwriter")
	if err != nil {
		t.Log(err)
	}
	writer.Flush()
}

// ioutil writefile 已弃用
func TestIoWriteFile(t *testing.T) {
	// file, err := os.Create("osIoWriteFile.txt")
	// if err != nil {
	// 	t.Log(err)
	// }
	content := []byte("hello io write file")
	err := ioutil.WriteFile("osIoWriteFile.txt", content, 0644)
	if err != nil {
		t.Log(err)
	}
}

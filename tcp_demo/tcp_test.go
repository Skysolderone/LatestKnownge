package test

import (
	"bytes"
	"encoding/binary"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestTcp(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		var err error
		_, err = conn.Write([]byte(strconv.Itoa(i) + "aaaa\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + "bbb\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + "cccc\n"))
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(time.Second)

}

func TestLengthClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		var err error
		data, err := Encode(strconv.Itoa(i) + "aaaa\n")
		_, err = conn.Write(data)
		data, err = Encode(strconv.Itoa(i) + "bbb\n")
		_, err = conn.Write(data)
		data, err = Encode(strconv.Itoa(i) + "cccc\n")
		_, err = conn.Write(data)
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(time.Second)
}
func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err

	}
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(reverse []byte) []byte {
	middle := len(reverse) / 2
	for i := 0; i != middle; i++ {
		reverse[i], reverse[len(reverse)-1-i] = reverse[len(reverse)-1-i], reverse[i]
	}
	return reverse
}

func TestRervers(t *testing.T) {
	initdata := []byte{'a', 's', 'd', 'f', 'g'}
	initdata2 := []byte{'a', 's', 'd', 'f'}
	file, _ := os.OpenFile("demo", os.O_RDWR, 0o644)
	defer file.Close()
	size, _ := file.Seek(0, io.SeekEnd)
	file.Seek(0, io.SeekStart)
	t.Run("write", func(t *testing.T) {
		reverse(initdata)
		reverse(initdata2)
		assert.Equal(t, []byte{'g', 'f', 'd', 's', 'a'}, initdata, "reverse failed")
		assert.Equal(t, []byte{'f', 'd', 's', 'a'}, initdata2, "reverse failed")
		file.WriteAt(initdata, size)
	})
	t.Run("read", func(t *testing.T) {
		buf := make([]byte, 1)
		data := make([]byte, 0, size)
		for i := 1; i <= int(size); i++ {
			file.Seek(-int64(i), io.SeekEnd)
			file.Read(buf)
			data = append(data, buf[0])
		}
		assert.Equal(t, []byte{'a', 's', 'd', 'f', 'g'}, data, "read reverse failed")
	})
}

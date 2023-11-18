package test

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	file, err := os.Open("./ostest.txt")
	if err != nil {
		t.Log(err)
	}
	defer file.Close()
}

package main

import (
	"bytes"
	"encoding/gob"
)

type Codec interface {
	Encoder(i any) ([]byte, error)
	Decoder(data []byte, i any) error
}
type Gobc struct{}

func (c *Gobc) Encoder(i any) ([]byte, error) {
	var buf bytes.Buffer
	encode := gob.NewEncoder(&buf)
	if err := encode.Encode(i); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (c *Gobc) Decoder(data []byte, i any) error {
	buf := bytes.NewBuffer(data)
	decode := gob.NewDecoder(buf)
	return decode.Decode(i)
}

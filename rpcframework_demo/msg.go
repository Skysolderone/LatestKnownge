package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type RPCMsg struct {
	*Header
	ServiceClass  string
	ServiceMethod string
	Payload       []byte
}

const SPLIT_LEN = 4

func NewRPC() *RPCMsg {
	header := Header([HEADER_LEN]byte{})
	header[0] = magicNumber
	return &RPCMsg{
		Header: &header,
	}
}
func (msg *RPCMsg) send(writer io.Writer) error {
	_, err := writer.Write(msg.Header[:])
	if err != nil {
		return err
	}
	dataLen := SPLIT_LEN + len(msg.ServiceClass) + SPLIT_LEN + len(msg.ServiceMethod) + SPLIT_LEN + len(msg.Payload)
	err = binary.Write(writer, binary.BigEndian, uint32(dataLen))
	if err != nil {
		return err
	}
	err = binary.Write(writer, binary.BigEndian, uint32(len(msg.ServiceClass)))
	if err != nil {
		return err
	}
	err = binary.Write(writer, binary.BigEndian, util.StringToByte(msg.ServiceClass))
	if err != nil {
		return err
	}
	return nil
}
func Read(r io.Reader) (*RPCMsg, error) {
	msg := NewRPC()
	err := msg.Decode(r)
	if err != nil {
		return nil, err
	}
	return msg, err

}
func (msg *RPCMsg) Decode(r io.Reader) error {
	_, err := io.ReadFull(r, msg.Header[:])
	if !msg.Header.CheckMagicNumber() {
		return fmt.Errorf("magic number error:%v\n", msg.Header[0])
	}
	headerByte := make([]byte, 4)
	_, err = io.ReadFull(r, headerByte)

	if err != nil {
		return err
	}
	bodyLen := binary.BigEndian.Uint32(headerByte)
	data := make([]byte, bodyLen)
	_, err = io.ReadFull(r, data)
	start := 0
	end := start + SPLIT_LEN
	classLen := binary.BigEndian.Uint32(data[start:end])
	start = end
	end = start + int(classLen)
	msg.ServiceClass = util.ByteToString(data[start:end])
}

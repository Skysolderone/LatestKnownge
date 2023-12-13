package main

const (
	HEADER_LEN = 5
)
const (
	magicNumber byte = 0x06
)

type MsgType byte

const (
	Request MsgType = iota
	Response
)

type CompressType byte

const (
	None CompressType = iota
	Gzip
)

type SerializeType byte

const (
	Gob SerializeType = iota
	JSON
)

type Header [HEADER_LEN]byte

func (h *Header) CheckMagicNumber() bool {
	return h[0] == magicNumber
}
func (h *Header) Version() byte {
	return h[1]
}
func (h *Header) SetVersion(version byte) {
	h[1] = version
}

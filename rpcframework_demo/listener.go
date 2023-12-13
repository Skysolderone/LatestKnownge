package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Listener interface {
	Run()
	SetHeader(string, Handler)
	Close()
}

type RPCListener struct {
	ServiceIp   string
	ServicePort int
	Handles     map[string]Handler
	nl          net.Listener
}

func NewRPCListener(serviceIP string, servicePort int) *RPCListener {
	return &RPCListener{
		ServiceIp:   serviceIP,
		ServicePort: servicePort,
		Handles:     make(map[string]Handler),
	}
}

func (l *RPCListener) Run() {
	addr := fmt.Sprintf("%s:%s", l.ServiceIp, l.ServicePort)
	nl, err := net.Listen(config.NET_TRANS_PROTOCOL, addr)
	if err != nil {
		panic(err)
	}
	l.nl = nl
	for {
		conn, err := l.nl.Accept()
		if err != nil {
			continue
		}
		go l.handelConn(conn)
	}
}
func (l *RPCListener) Close() {
	if l.nl != nil {
		l.nl.Close()
	}
}
func (l *RPCListener) SetHandler(name string, handler Handler) {
	if _, ok := l.Handlers[name]; ok {
		log.Printf("%s is registered!\n", name)
		return
	}
	l.Handlers[name] = handler
}
func (l *RPCListener) handleConn(conn net.Conn) {
	defer catchPanic()
	for {
		msg, err := l.receiveData(conn)
		if err != nil || msg == nil {
			return
		}
		coder := global.Codecs[msg.Header.SerializeType()]
		if coder == nil {
			return
		}
		inArgs := make([]interface{}, 0)
		err = coder.Decode(msg.Payload, &inArgs)
		if err != nil {
			return
		}
		handler, ok := l.Handlers[msg.ServiceClass]
		if !ok {
			return
		}
		result, err := handler.Handle(msg.ServiceMethod, inArgs)
		encodeRes, err := coder.Encode(result)
		if err != nil {
			return
		}
		err = l.sendData(conn, encodeRes)
		if err != nil {
			return
		}
	}
}
func (l *RPCListener) receiveData(conn net.Conn) (*protocol.RPCMsg, error) {
	msg, err := protocol.Read(conn)
	if err != nil {
		if err != io.EOF { //close
			return nil, err
		}
	}
	return msg, nil
}
func (l *RPCListener) sendData(conn net.Conn, payload []byte) error {
	resMsg := protocol.NewRPCMsg()
	resMsg.SetVersion(config.Protocol_MsgVersion)
	resMsg.SetMsgType(protocol.Response)
	resMsg.SetCompressType(protocol.None)
	resMsg.SetSerializeType(protocol.Gob)
	resMsg.Payload = payload
	return resMsg.Send(conn)
}

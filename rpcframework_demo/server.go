package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type Server interface {
	Register(string, any)
	Run()
	Close()
}

type RPCServer struct {
	listener Listener
}

func NewRPCServer(ip string, port int) *RPCServer {
	return &RPCServer{
		listener: NewRPCListener(ip, port),
	}
}

func (svr *RPCServer) Run() {
	go svr.listener.Run()
}
func (svr *RPCServer) Close() {
	if svr.listener != nil {
		svr.listener.Close()
	}
}
func (svr *RPCServer) Register(class interface{}) {
	name := reflect.Indirect(reflect.ValueOf(class)).Type().Name()
	svr.RegisterName(name, class)
}
func (svr *RPCServer) RegisterName(name string, class interface{}) {
	handler := &RPCServerHandler{class: reflect.ValueOf(class)}
	svr.listener.SetHandler(name, handler)
	log.Printf("%s registered success!\n", name)
}

type TestHandler struct{}

func (t *TestHandler) Hello() string {
	return "hello world"
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var userList = map[int]User{
	1: User{1, "hero", 11},
	2: User{2, "kavin", 12},
}

type UserHandler struct{}

func (u *UserHandler) GetUserById(id int) (User, error) {
	if u, ok := userList[id]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("id %d not found", id)
}
func main() {
	flag.Parse()
	if ip == "" || port == 0 {
		panic("init ip and port error")
	}
	srv := provider.NewRPCServer(ip, port)
	srv.RegisterName("User", &UserHandler{})
	srv.RegisterName("Test", &TestHandler{})
	gob.Register(User{})
	go srv.Run()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	srv.Close()
}

package helloapi

import (
	"context"
	hello "goademo/gen/hello"

	"goa.design/clue/log"
)

// hello service example implementation.
// The example methods log the requests and return zero values.
type hellosrvc struct{}

// NewHello returns the hello service implementation.
func NewHello() hello.Service {
	return &hellosrvc{}
}

// Say implements say.
func (s *hellosrvc) Say(ctx context.Context, p string) (res string, err error) {
	log.Printf(ctx, "hello.say")
	return
}

// Code generated by goa v3.20.1, DO NOT EDIT.
//
// hello service
//
// Command:
// $ goa gen goademo/design

package hello

import (
	"context"
)

// The hello service says hello
type Service interface {
	// Say implements say.
	Say(context.Context, string) (res string, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "hello"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hello"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"say"}

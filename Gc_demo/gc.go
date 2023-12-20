package main

import (
	"fmt"
	"runtime"
	"time"
)

// basic
func main() {
	// Enable GC to run manually for demonstration purposes
	runtime.GC()

	// Allocate a new object
	obj := &Object{name: "Sample Object"}

	// Create a reference to the object
	ref := obj
	fmt.Println(ref)
	// Set the reference to nil, making the original object unreachable
	ref = nil

	// Explicitly run the garbage collector to deallocate unreachable objects
	runtime.GC()

	// The program's output depends on GC behavior
	fmt.Println("Garbage collection example completed")
	go func() {
		fmt.Println("goexit ago")
		runtime.Goexit()
		fmt.Println("goexit after")
	}()
	time.Sleep(3 * time.Second)
}

type Object struct {
	name string
}

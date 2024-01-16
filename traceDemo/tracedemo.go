package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

// trace  追踪数据
func main() {
	f, err := os.Create("profile.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	traceFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()
	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		_ = square(n)

	}
}

func square(n int) int {
	return n * n
}

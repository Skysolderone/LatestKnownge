package main

import (
	"os"
	"runtime/pprof"
	"time"
)

// go tool pprof cpu.pprof
func main2() {
	cpufile, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	defer cpufile.Close()
	pprof.StartCPUProfile(cpufile)
	defer pprof.StopCPUProfile()
	for i := 0; i < 1000000; i++ {
		_ = i * i
	}
	time.Sleep(time.Second * 5)
}

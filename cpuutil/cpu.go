package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

func getCpuLoad() {
	loadInfo, _ := load.Avg()
	fmt.Printf("%#v\n", loadInfo)
}

func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range cpuInfos {
		fmt.Printf("%#v\n", v)
	}

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for range ticker.C {
			perenct, _ := cpu.Percent(time.Second, false)
			if len(perenct) == 0 {
				fmt.Println("get percent failed")
				return
			} else {
				fmt.Printf("cpu percent: %v\n", perenct[0])
			}
			fmt.Println(perenct)
		}
	}()
}

func getMemInfo() {
	v, _ := mem.VirtualMemory()
	fmt.Printf("%#v\n", v)
}

func getHostInfo() {
	v, _ := host.Info()
	fmt.Println(v)
}

// 获取网络io信息
func getNetIO() {
	netInfos, _ := net.IOCounters(true)
	for idx, netInfo := range netInfos {
		fmt.Printf("网络接口 %v: 发送字节: %v 接收字节: %v\n", idx, netInfo.BytesSent, netInfo.BytesRecv)
	}
}

// 获取本地ip
func getLocalIp() (string, error) {
	addrs, _ := net.Interfaces()

	for _, addr := range addrs {
		fmt.Printf("%#v\n", addr.Name)
		fmt.Printf("%#v\n", addr.Flags)
		fmt.Printf("%#v\n", addr.HardwareAddr)
		fmt.Printf("%#v\n", addr.MTU)
		fmt.Printf("%#v\n", addr.Name)
	}
	return "", fmt.Errorf("未找到本地IP")
}

// 获取pid列表
func getPids() {
	pids, err := net.Pids()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range pids {
		fmt.Printf("%#v\n", v)
	}
}

func connectpid() {
	stats, err := net.ConnectionsPid("", 475)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range stats {
		fmt.Printf("%#v\n", v)
	}
}

func getProcess() {
	var root []*process.Process
	processes, err := process.Processes()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range processes {
		// if v.Pid == 0 {
		// 	root, _ = v.Children()
		// 	break
		// }

		persnet, _ := v.CPUPercent()
		if persnet > 10 {
			fmt.Printf("pid%#v\n", v.Pid)
			name, _ := v.Name()

			fmt.Printf("name%#v\n", name)
			fmt.Println("cpu user:", persnet)
		}
	}
	fmt.Println(len(root))
	for _, v := range root {
		fmt.Printf("pid%#v\n", v.Pid)
		name, _ := v.Name()
		fmt.Printf("name%#v\n", name)
	}
}

func main() {
	// getCpuLoad()
	// getMemInfo()
	// getHostInfo()
	// getCpuInfo()
	// getNetIO()
	// getLocalIp()
	// getPids()
	// connectpid()
	getProcess()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Stop")
}

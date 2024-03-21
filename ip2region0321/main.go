package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var (
	dbPath string = "./ip2region.xdb"
	ipBuff []byte
)

func init() {
	var err error
	ipBuff, err = xdb.LoadContentFromFile(dbPath)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	searcher, err := xdb.NewWithBuffer(ipBuff)
	if err != nil {
		log.Fatal(err)
	}
	defer searcher.Close()
	ip := "146.190.85.111"
	startTime := time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("addr: %s, took: %s\n", region, time.Since(startTime))
}

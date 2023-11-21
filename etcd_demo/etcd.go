package main

import (
	"log"

	"github.com/zieckey/etcdsync"
)
//基于etcdsync 的分布式锁  You need a etcd instance running on http://localhost:2379, 需要docker etcd 
func main() {
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	err = m.lock()
	if err != nil {
		log.Println("etcdsync lock fail")
	}
	log.Println("etcdsync lock ok")
	log.Println("do something")
	err = m.unlock()
	if err != nil {
		log.Println("etcdsync unlock succ")
	} else {
		log.Println("etcdsync unlock fail")
	}

}

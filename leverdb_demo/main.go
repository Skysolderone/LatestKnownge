package main

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("db.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	value, err := db.Get([]byte("symbols"), nil)
	if err != nil {
		log.Fatal(err)
	}
	var s []string
	sonic.Unmarshal(value, &s)
	for _, v := range s {
		fmt.Println(string(v))
	}
	// s = append(s, "aevousdt")
	// res, _ := sonic.Marshal(s)
	// err = db.Put([]byte("symbols"), res, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(value))
	// symbols := []string{"btcusdt", "ethusdt", "solusdt", "ordiusdt"}
	// bys, err := sonic.Marshal(symbols)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// batch := new(leveldb.Batch)

	// batch.Put([]byte("symbols"), bys)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.Write(batch, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

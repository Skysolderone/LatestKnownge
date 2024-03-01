package main

import "log"

var CapitalMap = map[int]CapitalStr{
	1: {Capital: []int{500}, Auth: false},
	2: {Capital: []int{1000, 2000}, Auth: false},
	3: {Capital: []int{4000, 8000}, Auth: false},
	4: {Capital: []int{15000, 30000}, Auth: false},
	5: {Capital: []int{60000, 100000}, Auth: false},
	6: {Capital: []int{150000, 200000}, Auth: false},
}

type CapitalStr struct {
	Capital []int `json:"capital"`
	Auth    bool  `json:"auth"`
}

func main() {
	oldmap := CapitalMap
	for i := 1; i <= 3; i++ {
		authchildobj := CapitalStr{}
		res := CapitalMap[i]
		authchildobj.Capital = res.Capital
		authchildobj.Auth = true
		CapitalMap[i] = authchildobj

	}
	for _, v := range CapitalMap {
		log.Println(v)
	}
	CapitalMap = oldmap
	log.Println(CapitalMap)
}

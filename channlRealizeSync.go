package main

import "fmt"

type PerSync struct {
	ch chan interface{}
}

func NewPer() *PerSync {
	p := &PerSync{}
	p.ch = make(chan interface{}, 1)
	return p
}
func (p *PerSync) Lock() {
	p.ch <- struct{}{}
}
func (p *PerSync) Unlock() {
	<-p.ch
}
func main() {
	p := NewPer()
	count := 0
	p.Lock()
	count++
	//p.Lock()
	p.Unlock()
	fmt.Println(count)

}

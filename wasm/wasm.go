package main

import (
	"strconv"
	"syscall/js"
	"time"
)

var (
	document = js.Global().Get("document")
	numEle   = document.Call("getElementById", "num")
	ansEle   = document.Call("getElementById", "ans1")
	btnEle   = js.Global().Get("btn1")
)

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}
func fibFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(fib(args[0].Int()))
}
func fibFunc1(this js.Value, args []js.Value) interface{} {
	v := numEle.Get("value")
	if num, err := strconv.Atoi(v.String()); err != nil {
		ansEle.Set("innerHTML", js.ValueOf(fib(num)))
	}
	return nil
}
func fibFunc2(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1]
	go func() {
		time.Sleep(3 * time.Second)
		v := fib(args[0].Int())
		callback.Invoke(v)
	}()

	js.Global().Get("ans").Set("innerHTML", "Waiting 3s...")
	return nil
}
func main() {
	// alert := js.Global().Get("alert")
	// alert.Invoke("Hello world")

	done := make(chan int, 0)
	//js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	// btnEle.Call("addEventListener", "click", js.FuncOf(fibFunc1))
	js.Global().Set("fibFunc", js.FuncOf(fibFunc2))
	<-done

}

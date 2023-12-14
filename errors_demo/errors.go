package main

import (
	"fmt"

	"github.com/go-errors/errors"
)

var Crashed = errors.Errorf("test stackover")

func Crash() error {
	return errors.New(Crashed)
}

// errors库  获取错误显示的堆栈信息
func main1() {
	err := Crash()
	if err != nil {
		fmt.Println(err.(*errors.Error).ErrorStack())
		return
	}
}
func main() {
	err := crashy.Crash()
	if err != nil {
		if errors.Is(err, crashy.Crashed) {
			fmt.Println(err.(*errors.Error).ErrorStack())
		} else {
			panic(err)
		}
	}
}

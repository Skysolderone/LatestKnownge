package csv

import (
	"fmt"
	"os"
)

func Ge() {
	_, err := os.Create("../eee.csv")
	if err != nil {
		fmt.Println(err)
	}
	_, err = os.Create("./csv/aaa.csv")
	if err != nil {
		fmt.Println(err)
	}
}

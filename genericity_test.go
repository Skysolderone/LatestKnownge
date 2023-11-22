package test

import (
	"fmt"
	"testing"
)

func sum(n []interface{}) interface{} {
	var s float32
	for _, item := range n {
		switch t := item.(type) {
		case int32:
			s += float32(t)
		case float32:
			s += t
		}
	}
	if len(n) > 0 {
		if _, ok := n[0].(int32); ok {
			return int32(s)
		}
	}
	return s
}

type sumType interface {
	int32 | float32 | string
}

func sumGEN[T sumType](n []T) T {
	var s T
	for _, item := range n {

		s += item

	}
	return s
}
func TestGenericity(t *testing.T) {
	data1 := []int32{20, 50, 10, 30}
	data2 := []float32{32.3, 45.1, 54.3, 65.4}
	data3 := []string{"hello", "world"}
	// data1interface := make([]interface{}, len(data1))
	// data2interface := make([]interface{}, len(data2))
	// for i := range data1 {
	// 	data1interface[i] = data1[i]
	// }
	// for i := range data2 {
	// 	data2interface[i] = data2[i]
	// }
	// sum1 := sum(data1interface)
	// sum2 := sum(data2interface)
	//use genericity
	sum1 := sumGEN[int32](data1)
	sum2 := sumGEN[float32](data2)
	sum3 := sumGEN[string](data3)
	fmt.Printf("sum1:%v,(%T)\n", sum1, sum1)
	fmt.Printf("sum2:%v,(%T)\n", sum2, sum2)
	fmt.Printf("sum3:%v,(%T)\n", sum3, sum3)

}

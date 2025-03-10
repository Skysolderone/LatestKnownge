package main

import (
	"errors"
	"fmt"
	"sync"
)

func LocalModel[T comparable, key any, value comparable](T) {
}

type user struct {
	Name string
	Age  int
}
type book struct {
	Name string
	Age  int
}
type ModelType interface {
	user | book
}
type KeyModel interface {
	int | float64
}

var testm sync.Map

func load_model[T ModelType, key KeyModel](model T, k key, m *sync.Map) (T, error) {
	v, ok := m.Load(k)
	if !ok {
		return model, errors.New("not found")
	}
	return v.(T), nil
}

func main() {
	s1 := user{
		Name: "testuser",
		Age:  18,
	}
	s2 := book{
		Name: "testbook",
		Age:  20,
	}
	testm.Store(1, s1)
	// testm.Load(1)
	// testm.Delete(1)
	testm.Store(2, s2)
	t1, err := load_model(user{}, 1, &testm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1)
	t2, err := load_model(book{}, 2, &testm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	left  *Tree
	value int
	right *Tree
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.left)
	fmt.Println(t.value, " ")
	traverse(t.right)
}

func Create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.value {
		return t
	}
	if v < t.value {
		t.left = insert(t.left, v)
		return t
	}
	t.right = insert(t.right, v)
	return t
}
func main() {
	tree := Create(10)
	fmt.Println("tree value is:", tree.value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -1)
	traverse(tree)
	fmt.Println()
	fmt.Println("tree is :", tree.value)
}

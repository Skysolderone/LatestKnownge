package main

import (
	"errors"
	"log"
)

type Node[T any] struct {
	Value    []T
	Children map[rune]*Node[T]
}

func NewNode[T any]() *Node[T] {
	return &Node[T]{
		Children: make(map[rune]*Node[T], 0),
	}
}
func (n *Node[T]) getOrCreateChild(r rune) *Node[T] {
	if node, has := n.Children[r]; has {
		return node
	}
	node := NewNode[T]()
	n.Children[r] = node
	return node
}

var (
	ErrNotCompatibleKey = errors.New("the key is not compatible")
	ErrNotExists        = errors.New("the key not exists")
)

type Storage[T any] struct {
	root  *Node[T]
	depth int
}

func NewStorage[T any](depth int) *Storage[T] {
	return &Storage[T]{
		root:  NewNode[T](),
		depth: depth,
	}
}

func (t *Storage[T]) Add(key string, value T) error {
	if len(key) != t.depth {
		return ErrNotCompatibleKey
	}
	current := t.root
	for _, r := range key {
		current = current.getOrCreateChild(r)
	}
	current.Value = append(current.Value, value)
	return nil
}

func (t *Storage[T]) Search(key string) []T {
	cursor := t.root
	for _, r := range key {
		if node, has := cursor.Children[r]; !has {
			return nil
		} else {
			cursor = node
		}
	}
	return dfs(cursor)
}

func dfs[T any](node *Node[T]) []T {
	if len(node.Value) != 0 {
		return node.Value
	}
	var items []T
	for _, n := range node.Children {
		items = append(items, dfs(n)...)
	}
	return items
}
func main() {
	tr := NewStorage[int](3)
	tr.Add("hlm", 1)
	tr.Add("hlm", 2)
	tr.Add("hle", 3)
	tr.Add("hlf", 3)
	tr.Add("hlf", 5)
	tr.Add("hef", 5)
	tr.Add("hee", 4)
	tr.Add("eff", 10)
	tr.Add("efe", 12)
	tr.Add("ehe", 13)
	if l := len(tr.Search("ef")); l != 2 {
		log.Fatal("test failed for ef, expected 2, got ", l)
	}
	if l := len(tr.Search("hl")); l != 5 {
		log.Fatal("test failed for ef, expected 5, got ", l)
	}
	if l := len(tr.Search("e")); l != 3 {
		log.Fatal("test failed for ef, expected 3, got ", l)
	}
}

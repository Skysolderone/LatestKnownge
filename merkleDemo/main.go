package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// node
type Node struct {
	Left  *Node
	Right *Node
	Data  []byte
}

type Tree struct {
	Root *Node
}

func GetHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func NewNode(letf, right *Node, data []byte) *Node {
	mNode := Node{}
	if letf == nil && right == nil {
		mNode.Data = GetHash(data)
	} else {
		prehash := append(letf.Data, right.Data...)
		mNode.Data = GetHash(prehash)
	}
	mNode.Left = letf
	mNode.Right = right
	
	return &mNode
}

func NewTree(data [][]byte) *Tree {
	var nodes []Node
	for _, dataum := range data {
		node := NewNode(nil, nil, dataum)
		nodes = append(nodes, *node)
		fmt.Printf("node is :%s \n", hex.EncodeToString(node.Data))
	}

	for len(nodes) > 1 {
		var newLever []Node
		for i := 0; i < len(nodes); i += 2 {
			var left, right *Node
			left = &nodes[i]
			if i+1 < len(nodes) {
				right = &nodes[i+1]
			} else {
				right = &nodes[i]
			}
			newNode := NewNode(left, right, nil)
			newLever = append(newLever, *newNode)
			fmt.Printf("lever:%d\n", len(newLever))
			fmt.Printf("newnode is :%s \n", hex.EncodeToString(newNode.Data))

		}
		nodes = newLever
	}
	mTree := Tree{&nodes[0]}
	return &mTree
}

func verify(data []byte, path [][]byte, roothahs []byte) bool {
	currentHash := GetHash(data)
	for _, hash := range path {
		combind := append(currentHash, hash...)
		currentHash = GetHash(combind)
	}
	fmt.Printf("currentHash is :%s\n roothash is :%s\n", hex.EncodeToString(currentHash), hex.EncodeToString(roothahs))
	return hex.EncodeToString(currentHash) == hex.EncodeToString(roothahs)
}

func main() {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
		// []byte("e"),
	}
	tree := NewTree(data)
	fmt.Printf("Tree is :%s\n", hex.EncodeToString(tree.Root.Data))
	roothash := tree.Root.Data
	path := [][]byte{
		GetHash([]byte("b")),
		GetHash(append(GetHash([]byte("c")), GetHash([]byte("d"))...)),
	}
	isvalid := verify([]byte("a"), path, roothash)
	fmt.Printf("Is the data valid? %v\n", isvalid)

	fmt.Println(hex.EncodeToString(tree.Root.Left.Data))
	fmt.Println(hex.EncodeToString(tree.Root.Right.Data))
	// fmt.Println(hex.EncodeToString(GetHash([]byte("a"))))
}

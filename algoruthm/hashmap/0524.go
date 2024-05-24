package main

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strconv"
)

// 一致性hash
// HashRing represents the consistent hash ring
type HashRing struct {
	nodes       []int
	nodeMap     map[int]string
	virtualNode int
}

// NewHashRing creates a new HashRing
func NewHashRing(virtualNode int) *HashRing {
	return &HashRing{
		nodeMap:     make(map[int]string),
		virtualNode: virtualNode,
	}
}

// AddNode adds a node to the hash ring
func (h *HashRing) AddNode(node string) {
	for i := 0; i < h.virtualNode; i++ {
		virtualNodeKey := node + "#" + strconv.Itoa(i)
		hash := int(sha1Hash(virtualNodeKey))
		h.nodes = append(h.nodes, hash)
		h.nodeMap[hash] = node
	}
	sort.Ints(h.nodes)
}

// RemoveNode removes a node from the hash ring
func (h *HashRing) RemoveNode(node string) {
	for i := 0; i < h.virtualNode; i++ {
		virtualNodeKey := node + "#" + strconv.Itoa(i)
		hash := int(sha1Hash(virtualNodeKey))
		index := sort.SearchInts(h.nodes, hash)
		h.nodes = append(h.nodes[:index], h.nodes[index+1:]...)
		delete(h.nodeMap, hash)
	}
}

// GetNode returns the closest node in the hash ring for the given key
func (h *HashRing) GetNode(key string) string {
	hash := int(sha1Hash(key))
	index := sort.SearchInts(h.nodes, hash)
	if index >= len(h.nodes) {
		index = 0
	}
	return h.nodeMap[h.nodes[index]]
}

// sha1Hash returns a sha1 hash of a string
func sha1Hash(key string) uint32 {
	h := sha1.New()
	h.Write([]byte(key))
	bs := h.Sum(nil)
	return uint32((int(bs[0]) << 24) + (int(bs[1]) << 16) + (int(bs[2]) << 8) + int(bs[3]))
}

func main() {
	ring := NewHashRing(3)
	ring.AddNode("ServerA")
	ring.AddNode("ServerB")
	ring.AddNode("ServerC")
	keys := []string{"Key1", "Key2", "Key3", "Key4"}
	for _, key := range keys {
		fmt.Printf("Key %s is assigned to %s\n", key, ring.GetNode(key))
	}
	ring.RemoveNode("ServerB")
	fmt.Println("After removing ServerB")
	for _, key := range keys {
		fmt.Printf("Key %s is assigned to %s\n", key, ring.GetNode(key))
	}
}

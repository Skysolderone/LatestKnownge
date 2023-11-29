package main

import "math/rand"

type node struct {
	nexts      []*node
	key, value int
}
type SkipList struct {
	head *node
}

// Read
// 根据 key 读取 val，第二个 bool flag 反映 key 在 skiplist 中是否存在
func (s *SkipList) Get(key int) (int, bool) {
	// 根据 key 尝试检索对应的 node，如果 node 存在，则返回对应的 val
	if _node := s.search(key); _node != nil {
		return _node.value, true
	}
	return -1, false
}

// 从跳表中检索 key 对应的 node
func (s *SkipList) search(key int) *node {
	//from head
	move := s.head
	//from max heigt
	for level := len(s.head.nexts) - 1; level >= 0; level-- {
		//from right range
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}
		//if key==key  reutrn
		if move.nexts[level] != nil && move.nexts[level].key == key {
			return move.nexts[level]
		}
		//layer not found next layer
	}
	//all layre not found
	return nil
}

//Write

// roll 骰子，决定一个待插入的新节点在 skiplist 中最高层对应的 index
func (s *SkipList) roll() int {
	var level int
	//if it is >0  add 1
	for rand.Int() > 0 {
		level++
	}
	return level
}

// 将 key-val 对加入 skiplist
func (s *SkipList) Put(key, value int) {
	//if k v exist update kv  return
	if _node := s.search(key); _node != nil {
		_node.value = value
		return
	}
	//roll new node height
	level := s.roll()
	//level > max skiplist hegigt  algn
	for len(s.head.nexts)-1 < level {
		s.head.nexts = append(s.head.nexts, nil)
	}
	//create new node
	newNode := node{
		key: key, value: value, nexts: make([]*node, level+1),
	}
	//from head maxheigt
	move := s.head
	for level := level; level >= 0; level-- {
		//from right range
		for move.nexts[level] != nil && move.nexts[level].key < key {
			move = move.nexts[level]
		}
		//change pointer relaize
		newNode.nexts[level] = move.nexts[level]
		move.nexts[level] = &newNode
	}

}

//Delete



func main() {

}

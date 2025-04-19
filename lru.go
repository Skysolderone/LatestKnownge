package main

type Node struct {
	key   int
	value int
	next  *Node
	front *Node
}

type Lru struct {
	head    *Node
	end     *Node
	hashmap map[int]*Node
	limit   int
}

func NewLru(limit int) *Lru {
	lru := &Lru{limit: limit}
	lru.hashmap = make(map[int]*Node, limit)
}

func (l *Lru) gets(key int) int {
	if v, ok := l.hashmap[key]; ok {
		l.refresh(v)
		return v.value
	}
	return -1
}

func (l *Lru) refresh(s *Node) {
	if s == l.end {
		return
	}
	l.remove(s)
	l.add(s)
}

func (l *Lru) remove(s *Node) int {
	if s == l.end {
		l.end = l.end.front
		l.end.next = nil
	} else if s == l.head {
		l.head = l.head.next
		l.head.front = nil
	} else {
		s.front.next = s.next
		s.next.front = s.front
	}
	return s.key
}

func (l *Lru) add(v *Node) {
	if l.end != nil {
		l.end.next = v
		v.pre = l.end
		v.next = nil
	}
	l.end = v
	if l.head == nil {
		l.head = v
	}
}

func (l *Lru) puts(k, v int) {
	if v, ok := l.hashmap[key]; !ok {
		if len(l.hashmap) >= l.limit {
			oldkey := l.remove(v)
			delete(l.hashmap, oldkey)
		}
		node := Node{key: k, value: v}
		l.add(&node)
		l.hashmap[key] = &node

	} else {
		v.value = v
		l.refresh(v)
	}
}

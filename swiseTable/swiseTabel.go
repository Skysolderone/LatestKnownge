package main

import (
	"fmt"
	"time"
)

const (
	groupSize    = 16            // 每个组的大小
	empty        = 0             // 空槽位标记
	deleted      = 1             // 删除槽位标记
	metadataSize = groupSize / 8 // 每个组的元数据大小
)

type entry struct {
	key   string
	value interface{}
}

type SwissTable struct {
	metadata []byte  // 元数据数组
	entries  []entry // 存储键值对的数组
	size     int     // 当前存储的键值对数量
	capacity int     // 哈希表的总容量
}

func main() {
	// 标准库 map
	start := time.Now()
	m := make(map[string]interface{})
	for i := 0; i < 1000000; i++ {
		m[fmt.Sprintf("key%d", i)] = i
	}
	s := m["key0"]
	fmt.Println(s)
	fmt.Println("Standard map insert time:", time.Since(start))

	// Swiss Table
	start = time.Now()
	st := NewSwissTable()
	for i := 0; i < 1000000; i++ {
		st.Insert(fmt.Sprintf("key%d", i), i)
	}
	t, _ := st.find("key0")
	fmt.Println(t)
	fmt.Println("Swiss Table insert time:", time.Since(start))
}

func hash(key string) uint64 {
	h := uint64(5381)
	for i := 0; i < len(key); i++ {
		h = (h << 5) + h + uint64(key[i])
	}
	return h
}

func NewSwissTable() *SwissTable {
	return &SwissTable{
		metadata: make([]byte, groupSize/metadataSize),
		entries:  make([]entry, groupSize),
		size:     0,
		capacity: groupSize,
	}
}

func (st *SwissTable) find(key string) (int, bool) {
	h := hash(key)
	groupIndex := int(h % uint64(st.capacity/groupSize))
	start := groupIndex * groupSize

	for i := 0; i < groupSize; i++ {
		index := start + i
		if index >= st.capacity {
			index -= st.capacity
		}

		metadata := st.metadata[index/metadataSize]
		bit := byte(1 << (index % metadataSize))

		if metadata&bit == 0 {
			return -1, false // 未找到
		}

		if st.entries[index].key == key {
			return index, true // 找到
		}
	}

	return -1, false // 未找到
}

func (st *SwissTable) Insert(key string, value interface{}) {
	index, exists := st.find(key)
	if exists {
		st.entries[index].value = value
		return
	}

	if st.size >= st.capacity {
		st.resize()
	}

	h := hash(key)
	groupIndex := int(h % uint64(st.capacity/groupSize))
	start := groupIndex * groupSize

	for i := 0; i < groupSize; i++ {
		index := start + i
		if index >= st.capacity {
			index -= st.capacity
		}

		metadata := st.metadata[index/metadataSize]
		bit := byte(1 << (index % metadataSize))

		if metadata&bit == 0 {
			st.entries[index] = entry{key, value}
			st.metadata[index/metadataSize] |= bit
			st.size++
			return
		}
	}

	st.resize()
	st.Insert(key, value)
}

func (st *SwissTable) Delete(key string) {
	index, exists := st.find(key)
	if !exists {
		return
	}

	st.metadata[index/metadataSize] &^= byte(1 << (index % metadataSize))
	st.entries[index] = entry{"", nil}
	st.size--
}

func (st *SwissTable) resize() {
	newCapacity := st.capacity * 2
	newMetadata := make([]byte, newCapacity/metadataSize)
	newEntries := make([]entry, newCapacity)

	oldEntries := st.entries
	st.metadata = newMetadata
	st.entries = newEntries
	st.capacity = newCapacity
	st.size = 0

	for _, entry := range oldEntries {
		if entry.key != "" {
			st.Insert(entry.key, entry.value)
		}
	}
}

package main

// 堆是一个完全二叉树
// 堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值。
// 大顶堆： 堆中每一个节点的值都必须大于等于其子树中每个节点的值。
// 小顶堆：堆中每一个节点的值都必须小于等于其子树中每个节点的值。

type heap struct {
	m   []int
	len int
}

func buildHeap(s []int) *heap {
	n := len(s) - 1
	for i := n / 2; i > 0; i++ {
		heapf(s, n, i)
	}
	return &heap{s, n}

}
func heapf(m []int, n, i int) {
	for {
		maxPos := i
		if 2*i <= n && m[2*i] > m[i] {
			maxPos = 2 * i
		}
		if 2*i+1 <= n && m[2*i+1] > m[maxPos] {
			maxPos = 2*i + 1
		}
		if maxPos == 1 {
			break
		}
		m[i], m[maxPos] = m[maxPos], m[i]
		i = maxPos
	}
}
func (h *heap) push(data int) {
	h.len++
	h.m = append(h.m, data)
	i := h.len
	for i/2 > 0 && h.m[i/2] < h.m[i] {
		h.m[i/2], h.m[i] = h.m[i], h.m[i/2]
		i = i / 2
	}
}
func (h *heap) pop() int {
	if h.len < 1 {
		return -1
	}
	out := h.m[1]
	h.m[1] = h.m[h.len]
	h.len--
	heapf(h.m, h.len, 1)
	return out
}

// big heap
func main() {
	m := []int{0, 9, 3, 6, 2, 1, 7} //index 0 not set value
	h := buildHeap(m)               //init heap
	h.push(50)
	h.pop()
	fmt.Println(h.m)
}

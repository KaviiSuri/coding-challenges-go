package internal

import (
	"container/heap"
)

type HuffmanHeap []*Node

func NewHuffmanHeap() *HuffmanHeap {
	hh := &HuffmanHeap{}
	heap.Init(hh)
	return hh
}

func NewHuffmanHeapFromFreq(f Freqs) *HuffmanHeap {
	hh := NewHuffmanHeap()
	for c, w := range f {
		n := newNode(c, w)
		heap.Push(hh, &n)
	}
	return hh
}

func (hh HuffmanHeap) Len() int {
	return len(hh)
}

func (hh HuffmanHeap) Less(i, j int) bool {
	if hh[i].Weight == hh[j].Weight {
		return hh[i].Ch > hh[j].Ch
	}
	return hh[i].Weight > hh[j].Weight
}

func (hh HuffmanHeap) Swap(i, j int) {
	hh[i], hh[j] = hh[j], hh[i]
}
func (hh *HuffmanHeap) Push(x any) {
	item := x.(*Node)
	*hh = append(*hh, item)
}

func (hh *HuffmanHeap) Pop() any {
	old := *hh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*hh = old[0 : n-1]
	return item
}

func (hh *HuffmanHeap) BuildTree() *Node {
	var a, b, c *Node
	for len(*hh) > 1 {
		a = heap.Pop(hh).(*Node)
		b = heap.Pop(hh).(*Node)
		c = &Node{
			Left:   a,
			Right:  b,
			Weight: a.Weight + b.Weight,
		}
		heap.Push(hh, c)
	}
	if len(*hh) == 1 {
		return heap.Pop(hh).(*Node)
	}

	return c
}

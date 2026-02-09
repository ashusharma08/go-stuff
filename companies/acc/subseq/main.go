package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minSubseq([]int{2, 1, 3, 3}, 2))
	fmt.Println(minSubseq([]int{-1, -2, 3, 4}, 3))
	fmt.Println(minSubseq([]int{3, 4, 3, 3}, 2))
}

type HP struct {
	val int
	idx int
}

type pheap []*HP

func (h *pheap) Push(x any) {
	*h = append(*h, x.(*HP))
}
func (h *pheap) Pop() any {
	old := *h
	l := len(old)
	val := old[l-1]
	*h = old[:l-1]
	return val
}
func (h *pheap) Len() int {
	return len(*h)
}
func (h *pheap) Less(i, j int) bool {
	return (*h)[i].val < (*h)[j].val
}
func (h *pheap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func minSubseq(arr []int, size int) []int {
	php := &pheap{}
	heap.Init(php)
	for idx, item := range arr {
		heap.Push(php, &HP{val: item, idx: idx})
		if php.Len() > size {
			heap.Pop(php)
		}
	}
	result := make([]*HP, size)
	i := 0
	for range size {
		result[i] = heap.Pop(php).(*HP)
		i++
	}

	slices.SortFunc(result, func(x *HP, y *HP) int {
		return x.idx - y.idx
	})
	final := make([]int, size)
	for i, item := range result {
		final[i] = item.val
	}
	return final
}

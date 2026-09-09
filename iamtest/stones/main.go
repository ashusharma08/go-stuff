package main

import (
	"container/heap"
	"fmt"
)

type hp []int

func (h *hp) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() any {
	old := *h
	l := len(old)
	p := old[l-1]
	*h = old[:l-1]
	return p
}

func (h *hp) Less(x, y int) bool {
	return (*h)[x] > (*h)[y]
}

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Swap(x, y int) {
	(*h)[x], (*h)[y] = (*h)[y], (*h)[x]
}

func lastStoneWeight(stones []int) int {
	hhh := &hp{}
	heap.Init(hhh)
	for _, item := range stones {
		heap.Push(hhh, item)
	}
	for {
		if hhh.Len() == 0 {
			return 0
		}
		if hhh.Len() == 1 {
			return heap.Pop(hhh).(int)
		}
		max1 := heap.Pop(hhh)
		max2 := heap.Pop(hhh)
		diff := max1.(int) - max2.(int)
		if diff == 0 {
			continue
		}
		if diff < 0 {
			diff = -(diff)
		}
		heap.Push(hhh, diff)
	}
}
func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{1, 3}))
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

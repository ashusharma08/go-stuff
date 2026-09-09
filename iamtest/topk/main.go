package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topk(arr, 2))
}

type Counter struct {
	V int
	C int
}

type hh []Counter

func (h *hh) Push(x any) {
	*h = append(*h, x.(Counter))
}
func (h *hh) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
func (h *hh) Len() int {
	return len(*h)
}
func (h *hh) Less(i, j int) bool {
	return (*h)[i].C > (*h)[j].C
}
func (h *hh) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func topk(arr []int, k int) []int {
	mp := make(map[int]int)
	for _, item := range arr {
		mp[item]++
	}
	h := &hh{}
	heap.Init(h)
	for k, v := range mp {
		heap.Push(h, Counter{k, v})
	}
	res := make([]int, 0, k)
	for range k {
		v := heap.Pop(h)
		res = append(res, v.(Counter).V)
	}
	return res
}

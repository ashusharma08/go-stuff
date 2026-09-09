package main

import (
	"container/heap"
	"fmt"
)

func main() {
	// gas := []int{1, 2, 3, 4, 5}
	// cost := []int{3, 4, 5, 1, 2}
	// fmt.Println(getGasIndex(gas, cost))
	// fmt.Println(getGasIndex(
	// 	[]int{2, 3, 4},
	// 	[]int{3, 4, 3},
	// ))
	// fmt.Println(getGasIndex(
	// 	[]int{2, 3, 1, 1, 10},
	// 	[]int{1, 5, 2, 2, 1},
	// ))
	// fmt.Println(topk([]string{"Succession", "The Last of Us", "Succession", "Hacks", "The Last of Us", "Succession"}, 2))
	fmt.Println(isPossibleToStart(4, [][]int{{1, 0}, {2, 1}, {3, 2}}))
	fmt.Println(isPossibleToStart(4, [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}}))

}

func isPossibleToStart(num int, arr [][]int) bool {

	depend := make([][]int, num)
	indegree := make([]int, num)
	for _, item := range arr {
		depend[item[1]] = append(depend[item[1]], item[0])
		indegree[item[0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < num; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	level := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		level++

		for _, dependent := range depend[current] {
			indegree[dependent]--
			if indegree[dependent] == 0 {
				queue = append(queue, dependent)
			}
		}
	}
	fmt.Printf("%#v, %#v, %#v", depend, indegree, queue)
	return level == num
}

type gp struct {
	name  string
	count int
}
type hp []*gp

func (h *hp) Push(x any) {
	whatever := x.(*gp)
	(*h) = append((*h), whatever)
}
func (h *hp) Pop() any {
	l := len(*h)
	p := (*h)[l-1]
	*h = (*h)[:l-1]
	return p
}
func (h *hp) Len() int {
	return len(*h)
}
func (h *hp) Less(i, j int) bool {
	return (*h)[i].count > (*h)[j].count
}
func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func topk(movies []string, k int) []string {
	h := &hp{}
	heap.Init(h)

	mp := make(map[string]int, 0)
	for _, item := range movies {
		mp[item]++
	}

	for k, v := range mp {
		heap.Push(h, &gp{name: k, count: v})
	}
	res := make([]string, 0, k)
	for range k {
		res = append(res, heap.Pop(h).(*gp).name)
	}
	return res
}

func getGasIndex(gas []int, cost []int) int {
	start := 0
	curr, overall := 0, 0
	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		curr += diff
		overall += diff
		if diff < 0 {
			curr = 0
			start = i + 1
		}
	}

	if overall < 0 {
		return -1
	}

	return start
}

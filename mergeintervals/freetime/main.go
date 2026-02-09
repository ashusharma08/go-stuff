package main

import (
	"container/heap"
	"fmt"
)

func main() {
	schedule := [][]Interval{
		{{1, 3}, {5, 6}},
		{{2, 3}, {6, 8}},
	}
	freeTime := findEmployeeFreeTime(schedule)
	fmt.Println("Common free time slots:", freeTime)
}

type Interval struct {
	Start, End int
}
type Entry struct {
	interval Interval
	EmpIdx   int
	intIdx   int
}

type minHeap []Entry

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i, j int) bool {
	return m[i].interval.Start < m[j].interval.Start
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(Entry))
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func findEmployeeFreeTime(schedule [][]Interval) []Interval {
	hp := &minHeap{}
	heap.Init(hp)
	for i := range schedule {
		if len(schedule[i]) > 0 {
			heap.Push(hp, Entry{schedule[i][0], i, 0})
		}
	}
	res := make([]Interval, 0)
	prev := (*hp)[0].interval
	for hp.Len() > 0 {
		ent := heap.Pop(hp).(Entry)
		curr := ent.interval
		if prev.End < curr.Start {
			res = append(res, Interval{prev.End, curr.Start})
		}
		if curr.End > prev.End {
			prev = curr
		}
		if ent.intIdx+1 < len(schedule[ent.EmpIdx]) {
			nextInt := schedule[ent.EmpIdx][ent.intIdx+1]
			heap.Push(hp, Entry{nextInt, ent.EmpIdx, ent.intIdx + 1})
		}
	}
	return res
}

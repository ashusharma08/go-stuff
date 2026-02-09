package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Job struct {
	Start   int
	End     int
	CPULoad int
}

func main() {
	jobs1 := []*Job{{Start: 1, End: 4, CPULoad: 3}, {Start: 2, End: 5, CPULoad: 4}, {Start: 7, End: 9, CPULoad: 6}}
	fmt.Printf("Maximum CPU load at any time: %d\n", findMaxCPULoad(jobs1))

	jobs2 := []*Job{{Start: 6, End: 7, CPULoad: 10}, {Start: 2, End: 4, CPULoad: 11}, {Start: 8, End: 12, CPULoad: 15}}
	fmt.Printf("Maximum CPU load at any time: %d\n", findMaxCPULoad(jobs2))

	jobs3 := []*Job{{Start: 1, End: 4, CPULoad: 2}, {Start: 2, End: 4, CPULoad: 1}, {Start: 3, End: 6, CPULoad: 5}}
	fmt.Printf("Maximum CPU load at any time: %d\n", findMaxCPULoad(jobs3))
}

type jobHeap []*Job

func (j *jobHeap) Push(x any) {
	*j = append(*j, x.(*Job))
}
func (j *jobHeap) Pop() any {
	old := *j
	n := len(old)
	x := old[n-1]
	*j = old[:n-1]
	return x
}
func (j *jobHeap) Len() int {
	return len(*j)
}
func (j jobHeap) Less(i, k int) bool {
	return j[i].End < j[k].End
}
func (j *jobHeap) Swap(i, k int) {
	(*j)[i], (*j)[k] = (*j)[k], (*j)[i]
}

func findMaxCPULoad(jobs []*Job) int {
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Start < jobs[j].Start
	})
	maxCPULoad := 0
	currentCPULoad := 0
	jh := &jobHeap{}
	heap.Init(jh)
	for _, item := range jobs {
		for jh.Len() > 0 && item.Start > (*jh)[0].End {
			p := heap.Pop(jh).(*Job)
			currentCPULoad -= p.CPULoad
		}
		heap.Push(jh, item)
		currentCPULoad += item.CPULoad
		maxCPULoad = int(math.Max(float64(currentCPULoad), float64(maxCPULoad)))

	}
	return maxCPULoad
}

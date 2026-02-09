package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Meeting struct {
	Start int
	End   int
}

func main() {
	input := []Meeting{{1, 4}, {2, 5}, {7, 9}}
	result := findMinimumMeetingRooms(input)
	fmt.Printf("Minimum meeting rooms required: %d\n", result)

	input = []Meeting{{6, 7}, {2, 4}, {8, 12}}
	result = findMinimumMeetingRooms(input)
	fmt.Printf("Minimum meeting rooms required: %d\n", result)

	input = []Meeting{{1, 4}, {2, 3}, {3, 6}}
	result = findMinimumMeetingRooms(input)
	fmt.Printf("Minimum meeting rooms required: %d\n", result)

	input = []Meeting{{4, 5}, {2, 3}, {2, 4}, {3, 5}}
	result = findMinimumMeetingRooms(input)
	fmt.Printf("Minimum meeting rooms required: %d\n", result)
}

type meetingHeap []Meeting

func (m *meetingHeap) Push(x interface{}) {
	*m = append(*m, x.(Meeting))
}
func (m *meetingHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func (m meetingHeap) Len() int {
	return len(m)
}
func (m meetingHeap) Less(i, j int) bool {
	return m[i].End < m[j].End
}
func (m meetingHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func findMinimumMeetingRooms(meetings []Meeting) (mrooms int) {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].Start < meetings[j].Start
	})
	minHeap := &meetingHeap{}
	heap.Init(minHeap)
	for _, item := range meetings {
		for minHeap.Len() > 0 && item.Start >= (*minHeap)[0].End {
			heap.Pop(minHeap)
		}
		heap.Push(minHeap, item)
		if minHeap.Len() > mrooms {
			mrooms = minHeap.Len()
		}
	}
	return
}

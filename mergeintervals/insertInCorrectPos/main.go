package main

import (
	"fmt"
	"math"
)

type Interval struct {
	Start, End int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	newInt := make([]Interval, 0)
	i := 0
	for i < len(intervals) && newInterval.Start > intervals[i].End {
		newInt = append(newInt, intervals[i])
		i++
	}
	for i < len(intervals) && newInterval.End >= intervals[i].Start {
		newInterval.Start = int(math.Min(float64(intervals[i].Start), float64(newInterval.Start)))
		newInterval.End = int(math.Max(float64(intervals[i].End), float64(newInterval.End)))
		i++
	}
	newInt = append(newInt, newInterval)
	for i < len(intervals) {
		newInt = append(newInt, intervals[i])
		i++
	}
	return newInt
}

func main() {
	input := []Interval{{1, 3}, {5, 7}, {8, 12}}
	fmt.Print("Intervals after inserting the new interval: ")
	for _, interval := range insert(input, Interval{4, 6}) {
		fmt.Printf("[%d,%d] ", interval.Start, interval.End)
	}
	fmt.Println()

	input = []Interval{{1, 3}, {5, 7}, {8, 12}}
	fmt.Print("Intervals after inserting the new interval: ")
	for _, interval := range insert(input, Interval{4, 10}) {
		fmt.Printf("[%d,%d] ", interval.Start, interval.End)
	}
	fmt.Println()

	input = []Interval{{2, 3}, {5, 7}}
	fmt.Print("Intervals after inserting the new interval: ")
	for _, interval := range insert(input, Interval{1, 4}) {
		fmt.Printf("[%d,%d] ", interval.Start, interval.End)
	}
	fmt.Println()
}

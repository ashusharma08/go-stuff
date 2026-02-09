package main

import (
	"fmt"
	"math"
)

type Interval struct {
	Start, End int
}

func main() {
	input1 := []Interval{{Start: 1, End: 3}, {Start: 5, End: 6}, {Start: 7, End: 9}}
	input2 := []Interval{{Start: 2, End: 3}, {Start: 5, End: 7}}
	result := merge(input1, input2)
	fmt.Print("Intervals Intersection: ")
	for _, interval := range result {
		fmt.Printf("[%d,%d] ", interval.Start, interval.End)
	}
	fmt.Println()

}
func merge(arr1, arr2 []Interval) []Interval {
	i, j := 0, 0
	mergedIntervals := make([]Interval, 0)
	for i < len(arr1) && j < len(arr2) {
		if (arr1[i].Start >= arr2[j].Start && arr1[i].Start <= arr2[j].End) || (arr2[j].Start >= arr1[i].Start && arr2[j].Start <= arr1[i].End) {
			mergedIntervals = append(mergedIntervals, Interval{Start: int(math.Max(float64(arr1[i].Start), float64(arr2[j].Start))), End: int(math.Min(float64(arr1[i].End), float64(arr2[j].End)))})
		}
		if arr1[i].End < arr2[j].End {
			i++
		} else {
			j++
		}
	}
	return mergedIntervals
}

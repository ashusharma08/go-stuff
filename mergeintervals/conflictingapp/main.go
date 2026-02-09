package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func canAttendAllAppointments(intervals []Interval) bool {
	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x].Start < intervals[y].Start
	})
	i, j := 0, 1
	for j < len(intervals) {
		if intervals[i].End > intervals[j].Start {
			return false
		}
		i += 1
		j += 1
	}
	return true
}

func main() {
	intervals := []Interval{{Start: 1, End: 4}, {Start: 2, End: 5}, {Start: 7, End: 9}}
	result := canAttendAllAppointments(intervals)
	fmt.Println("Can attend all appointments 1: ", result)

	intervals1 := []Interval{{Start: 6, End: 7}, {Start: 2, End: 4}, {Start: 8, End: 12}}
	result = canAttendAllAppointments(intervals1)
	fmt.Println("Can attend all appointments 2 : ", result)

	intervals2 := []Interval{{Start: 4, End: 5}, {Start: 2, End: 3}, {Start: 3, End: 6}}
	result = canAttendAllAppointments(intervals2)
	fmt.Println("Can attend all appointments 3: ", result)
}

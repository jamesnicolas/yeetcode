package main

import (
	"fmt"
	"sort"
)

type intervalCol struct {
	intervals [][]int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (in intervalCol) Less(i, j int) bool {
	return in.intervals[i][0] < in.intervals[j][0]
}
func (in intervalCol) Len() int {
	return len(in.intervals)
}
func (in intervalCol) Swap(i, j int) {
	in.intervals[i], in.intervals[j] = in.intervals[j], in.intervals[i]
}

func mergeTwo(a, b []int) [][]int {
	if a[0] > b[0] {
		return mergeTwo(b, a)
	}
	if a[0] == b[0] {
		a[1] = max(a[1], b[1])
		return [][]int{a}
	}
	if a[1] < b[0] {
		return [][]int{a, b}
	}
	if a[1] >= b[0] {
		a[1] = max(a[1], b[1])
		return [][]int{a}
	}
	return [][]int{a, b}
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	ins := intervalCol{intervals}
	sort.Sort(ins)
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for _, v := range intervals[1:] {
		lastElementIndex := len(merged) - 1
		merged = append(merged[:lastElementIndex], mergeTwo(merged[lastElementIndex], v)...)
	}
	return merged
}

func main() {

	a := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	b := merge(a)
	fmt.Println(b)
}

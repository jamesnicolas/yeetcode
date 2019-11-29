package main

import (
	"fmt"
	"sort"
)

type intervalCol struct {
	intervals [][]int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (in intervalCol) Less(i, j int) bool {
	if in.intervals[i][0] == in.intervals[j][0] {
		return in.intervals[i][1] < in.intervals[j][1]
	}
	return in.intervals[i][0] < in.intervals[j][0]
}
func (in intervalCol) Len() int {
	return len(in.intervals)
}
func (in intervalCol) Swap(i, j int) {
	in.intervals[i], in.intervals[j] = in.intervals[j], in.intervals[i]
}

func mergeTwo(a, b []int) []int {
}

func merge(intervals [][]int) [][]int {
	ins := intervalCol{intervals}
	sort.Sort(ins)
	return ins.intervals
}

func main() {
	fmt.Println("asdf")
}

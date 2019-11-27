package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	pairs := make([][2]int, 0, len(nums)*(len(nums)-1)/2)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			pair := [2]int{i, j}
			pairs = append(pairs, pair)
		}
	}
	m := make(map[int][][2]int)
	solutions := make(map[[4]int]struct{})
	for _, v := range pairs {
		m[nums[v[0]]+nums[v[1]]] = append(m[nums[v[0]]+nums[v[1]]], [2]int{v[0], v[1]})
		if w, ok := m[target-(nums[v[0]]+nums[v[1]])]; ok {
			for _, x := range w {
				if x[0] != v[0] && x[0] != v[1] && x[1] != v[0] && x[1] != v[1] {
					solution := [4]int{nums[v[0]], nums[v[1]], nums[x[0]], nums[x[1]]}
					sort.Ints(solution[0:4])
					solutions[solution] = struct{}{}
				}
			}
		}
	}
	solutionsSlice := make([][]int, 0)
	for i := range solutions {
		sol := make([]int, 4)
		copy(sol, i[:])
		solutionsSlice = append(solutionsSlice, sol)
	}
	return solutionsSlice
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{0, 0}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(fourSum([]int{2, -4, -5, -2, -3, -5, 0, 4, -2}, -14))
}

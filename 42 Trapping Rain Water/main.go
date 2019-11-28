package main

import (
	"fmt"
)

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
func trap1(height []int) int {
	waterVol := 0
	for i, h := range height {
		maxR := 0
		for _, r := range height[i:] {
			maxR = max(maxR, r)
		}
		maxL := 0
		for _, l := range height[:i+1] {
			maxL = max(maxL, l)
		}

		waterVol += min(maxR, maxL) - h
		//fmt.Printf("maxR: %d, maxL %d, waterHeight %d\n", maxR, maxL, min(maxR,maxL) - h)
	}
	return waterVol
}

func main() {
	fmt.Println(trap1([]int{0, 1, 0, 1, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

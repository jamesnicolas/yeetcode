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

// brute force solution
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

// dp solution
func trap2(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	maxL := make([]int, length)
	maxR := make([]int, length)
	maxL[0] = height[0]
	maxR[length-1] = height[length-1]
	for i := length - 2; i > 0; i-- {
		maxR[i] = max(height[i], maxR[i+1])
	}
	waterVol := 0
	for i := 1; i < length; i++ {
		maxL[i] = max(height[i], maxL[i-1])
		waterVol += min(maxR[i], maxL[i]) - height[i]
	}
	return waterVol
}

func main() {
	fmt.Println(trap2([]int{0, 1, 0, 1, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	if length == 1 {
		return []int{1}
	}
	answer := make([]int, length, length)
	leftProduct := make([]int, length, length)
	rightProduct := make([]int, length, length)
	rightProduct[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}
	leftProduct[0] = 1
	answer[0] = rightProduct[0]
	for i := 1; i < length; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
		answer[i] = rightProduct[i] * leftProduct[i]
	}
	return answer
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

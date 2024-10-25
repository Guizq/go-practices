package main

import (
	"fmt"
	"math"
)

func main() {
	nums := make([]int, 50)
	for i := 0; i < 50; i++ {
		nums[i] = int(math.Floor(float64(i / 10)))
	}
	remval := 3
	removeVal(nums, remval)
	fmt.Println(nums)
}

func removeVal(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == val {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}
	return slow
}

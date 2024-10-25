package main

import (
	"fmt"
	"math"
)

func main() {
	var numbers []int = make([]int, 30)
	for i := 1; i < len(numbers)+1; i++ {
		numbers[i-1] = i*int(math.Pow(float64(-1), float64(i))) - 8
	}
	small := movalwin(numbers, 140)
	fmt.Println(small)
}

func movalwin(nums []int, target int) int {
	smallestdiffer := len(nums)
	differ := 0
	j2 := 0
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}
	for i := 0; i < len(nums); i++ {
		j := i
		sum := 0
		for sum < target && j < len(nums) {
			sum += nums[j]
			j2 = j
			j++
		}
		differ = j2 - i
		if differ < smallestdiffer && sum >= target {
			smallestdiffer = differ
		}
		fmt.Println(smallestdiffer)
	}
	return smallestdiffer + 1
}

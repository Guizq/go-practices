package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := make([]int, 10)
	for i := 1; i < 11; i++ {
		nums[i-1] = i*int(math.Pow(float64(-1), float64(i))) - 8
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)
}

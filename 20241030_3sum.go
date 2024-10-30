package main

import (
	"fmt"
	"sort"
)

func m332n() {
	sum3 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threesum(sum3))

}
func threesum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			n2 := nums[left]
			n3 := nums[right]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for left < right && n2 == nums[left] {
					left++
				}
				for left < right && n3 == nums[right] {
					right--
				}
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return res
}

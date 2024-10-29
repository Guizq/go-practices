package main

import "fmt"

func merere() {
	nums := []int{1, 2, 3, 4, 5}
	target := 8
	group := make([]int, 2)
	group = search2(nums, target)
	fmt.Println(group)
}
func search2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if preindex, ok := m[target-v]; ok {
			return []int{preindex, i}
		} else {
			m[v] = i
		}
	}
	return []int{}
}

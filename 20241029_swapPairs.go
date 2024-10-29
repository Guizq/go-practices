package main

import "fmt"

func intersection(num1 []int, num2 []int) []int {
	set := make(map[int]struct{}, 0)
	res := make([]int, 0)
	for _, v := range num1 {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
		}
	}
	for _, v := range num2 {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}
	return res
}

func m1323() {
	nums1 := []int{1, 2, 2, 3}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

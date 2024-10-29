package main

import "fmt"

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 0}
	D := []int{2, 0}
	fmt.Println(forinsu(A, B, C, D))
}
func forinsu(A []int, B []int, C []int, D []int) int {
	m1 := make(map[int]int)
	for _, v := range A {
		for _, u := range B {
			m1[u+v]++
		}
	}
	count := 0
	for _, v := range C {
		for _, u := range D {
			if preindex, ok := m1[-u-v]; ok {
				count += preindex
			}
		}
	}
	return count
}

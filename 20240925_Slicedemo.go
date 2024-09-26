package main

import "fmt"

func main() {
	//	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//		s1 := arr[2:6]
	//	s2 := arr[3:5]
	//	s3 := arr[:3]
	//	fmt.Println(s1, s2)
	//	fmt.Printf("the longth of the s1 is %d , cap is %d ", len(s1), cap(s1))
	//	fmt.Printf("%p\n", s1)
	//	fmt.Printf("%p\n", &arr[2])
	//	fmt.Printf("%p\n", &arr[0])
	//	fmt.Printf("%p\n", &arr)

	//	s3 = append(s3, 1, 2, 87, 34)
	//	fmt.Println(s3)
	//	fmt.Println(arr)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := make([]int, 0)
	for i, _ := range s1 {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1, s2)

	copy(s2, s1)
}

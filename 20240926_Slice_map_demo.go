package main

import "fmt"

func main() {
	var map1 map[int]string
	map1 = make(map[int]string)
	map1[100] = "wtf"
	map1[200] = "man"

	fmt.Println(map1)
	fmt.Println(map1[100])
	fmt.Println(map1[11]) //不存在 默认值string
	//map中没有index下
	value, ok := map1[1]
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
	fmt.Println(value)

	var map2 = map[int]float64{32: 98, 30: 9}
	for i, v := range map2 {
		fmt.Println(i, v)
	}

	user1 := map[string]string{"you": "are", "powerless": "!"}
	fmt.Println(user1["you"])
}

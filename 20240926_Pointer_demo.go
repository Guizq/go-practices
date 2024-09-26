package main

import "fmt"

func main() {
	ptr := why()
	fmt.Println(ptr)
	a := 10
	f6(&a)
	fmt.Println(a)
	fmt.Println(&a)
	type man struct {
		you   string
		power string
	}
	what := man{you: "are", power: "less"}
	fmt.Println(what)
	what.you = "wtf"
	fmt.Println(what)
	type address struct {
		city, xiaoqu string
	}
	type man struct {
		name string
		age  int
		addr address
	}
}
func why() *[6]int {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	return &arr
}
func f6(ptr *int) {
	*ptr = 20
}

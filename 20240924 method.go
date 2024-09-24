package main

import "fmt"

type rectangle struct{ width, height float64 }
type triangle struct{ width, height float64 }

func main() {
	r1 := rectangle{10, 40}
	t1 := triangle{20, 70}
	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", t1.area())
}
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (t triangle) area() float64 {
	return t.width * t.height * 0.5
}

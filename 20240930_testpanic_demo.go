package main

import "fmt"

func main() {
	defer fmt.Println("main1")
	defer fmt.Println("main2")
	testpanic(1)
	defer fmt.Println("main4")
	defer fmt.Println("main5")
}

func testpanic(i int) {
	defer fmt.Println("testmain1")
	defer fmt.Println("testmain2")
	if i == 1 {
		panic("bug!!!!")
	}
	defer fmt.Println("testmain4")
	defer fmt.Println("testmain5")
}

package main

import "fmt"

func main() {
	i := I{}
	assertsString(true)

	testAsserts(i)
}

type I interface {
}

func testAsserts(i interface{}) {
	switch v := i.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case I:
		fmt.Println("I")
	case interface{}:
		fmt.Println("interface")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("i don' know")
	}
}
func assertsString(i interface{}) {
	//断言失败则panic，程序停止运行
	s := i.(string)
	fmt.Println(s)
}

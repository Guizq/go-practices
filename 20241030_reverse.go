package main

import "fmt"

func main() {
	var strbyte []byte
	fmt.Scanln(&strbyte)
	for i := 0; i < len(strbyte); i++ {
		n1 := strbyte[i]
		if n1 >= '0' && n1 <= '9' {
			insert := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strbyte = append(strbyte[:i], append(insert, strbyte[i+1:]...)...)
			i += len(insert) - 1
		}
	}
	fmt.Println(string(strbyte))
}

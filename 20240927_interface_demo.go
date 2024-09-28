package main

import "fmt"

type USB interface {
	input()
	output()
}

func test(interfacetest USB) {
	interfacetest.input()
	interfacetest.output()
}

type Mouse struct {
	name string
}

func (mouse Mouse) input()  { fmt.Println("Mouse input") }
func (mouse Mouse) output() { fmt.Println("Mouse output") }

type Keyboard struct {
	name string
}

func (Keyboard Keyboard) input()  { fmt.Println("Keyboard input") }
func (Keyboard Keyboard) output() { fmt.Println("Keyboard output") }

func main() {
	keyboard1 := Keyboard{name: "wooting"}
	test(keyboard1)
	mouse1 := Mouse{name: "logitech"}
	test(mouse1)
}

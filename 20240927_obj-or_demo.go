package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	Person
	school string
}

type Valoranthero struct {
	name    string
	guntype string
}

type Fighter struct {
	Valoranthero
	dashskill string
}

type Assister struct {
	Valoranthero
	saveskill string
}

func (hero Valoranthero) move() {
	fmt.Println(hero, "is shooting")
}

func (fighter Fighter) dash() {
	fmt.Println(fighter, "dash to spite")
}

func (assister Assister) save() {
	fmt.Println(assister, "save fighter")
}

func main() {
	gui := Person{name: "kevin", age: 18}
	gui2 := Student{Person: Person{name: "halen", age: 20}, school: "xmu"}
	fmt.Println(gui)
	fmt.Println(gui2)
	fmt.Println(gui2.name)

	jettking := Fighter{Valoranthero{name: "jett", guntype: "shortshot"}, "dash"}
	sageking := Assister{Valoranthero{name: "sage", guntype: "longshot"}, "save"}

	jettking.move()
	jettking.dash()
	sageking.save()

	fmt.Printf("%s\t", jettking.name)
	fmt.Println("win the game")
}

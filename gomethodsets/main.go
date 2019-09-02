package main

import (
	"fmt"
)

type person struct {
	first string
	age int
}

func (p *person) speak() {
	fmt.Println(p.first, "aged", p.age, "says Hello!!!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first:"James",
		age : 30,
	}
	saySomething(&p)
	//Following generates error
	//saySomething(p)
}
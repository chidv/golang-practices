package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func changeMe (p1 person, p2 *person) {
	p1.age++
	p2.age++
}

func main() {
	x := 42
	fmt.Println(x, &x)

	//Pointer - pass by reference
	p1 := person { first: "James", last: "Bond", age:30}
	p2 := person { first: "Sherlock", last: "Holmes", age:32}

	fmt.Println("Before Changing")
	fmt.Println(p1)
	fmt.Println(p2)

	changeMe(p1, &p2)

	fmt.Println("After Changing")
	fmt.Println(p1)
	fmt.Println(p2)
}
package main

import (
	"fmt"
	"strings"
	"math"
)

type person struct {
	last string
	first string
	age int
}

type square struct {
	length float64
	width float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius*c.radius*math.Pi
}

func (s square) area() float64 {
	return s.length*s.width
}

//Polymorphism
func INFO(s shape) {
	fmt.Printf("Area of the shape %T is %f\n", s, s.area())
}

func (p person) speak() {
	fmt.Println("Hi")
	fmt.Println(strings.Title(strings.ToLower(p.first)), 
				strings.Title(strings.ToLower(p.last)), 
				"is", p.age, "years old")
}

func bar() (int, string) {
	return 1, "Hello World"
}

func foo() int {
	return 2
}

func foovariadic(input ...int) int {
	defer fmt.Println("Exiting the function foovariadic")
	fmt.Println("Entering the function foovariadic")
	fmt.Printf("%T\n", input)
	sum := 0
	for _, v := range input {
		sum = sum + v
	}
	return sum
}

func barslice(input []int) int {
	//defer example
	defer fmt.Println("Exiting the function barslice")
	fmt.Println("Entering the function barslice")

	fmt.Printf("%T\n", input)
	sum := 0
	for _, v := range input {
		sum = sum + v
	}
	return sum
}

func returnFunc(x int) func() int {
	//Closures
	f := func() int {
		x++
		return x
	}
	return f
}

func callbackFunc(fu func() int) {
	fmt.Println("Callback Function")
	fmt.Println(fu())
}

func main() {
	fr := foo()
	bri, brs := bar()

	fmt.Println(fr)
	fmt.Println(bri, brs)

	fmt.Println("Sum in main is", foovariadic(1,2,3,4,5,6,7,8,9,10))
	fmt.Println("Sum in main is", foovariadic(11,12,13,14,15,16))

	fmt.Println("Sum in main is", barslice([]int{1,2,3,4,5,6,7,8,9,10}))
	fmt.Println("Sum in main is", barslice([]int{1,2,3,4,5}))


	//Method Binding
	p := person{last:"Bond", first:"James", age:30}
	p.speak()

	//Polymorphism
	c := circle{radius: 23.5}
	s := square{length:10.5, width:23.7}

	INFO(c)
	INFO(s)

	//Anonymous Function
	func (input int) {
		fmt.Println("In Anonymous Function with input", input)
	}(77)


	//Function storing in a variable
	f := func (input int) {
		fmt.Println("In Assigned function with input", input)
	}
	f(78)

	//Return and Closures Func
	fu := returnFunc(77)
	fmt.Println(fu())
	fmt.Println(fu())

	//Callback functions
	callbackFunc(fu)
}
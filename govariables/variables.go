package main

import "fmt"

//Global Variables declared using var keyword
var gx int
var gy string
var gz bool

//Global Variables declared and initialized using var keyword
var x2 int= 42
var y2 string = "James Bond"
var z2 bool = true

//Userdefined type

type myType int

//myType variable declaration
var myVar myType
var myNewVar int


const cx int = 43

const (
	ucx = 44.1
)

const (
	Y20 = iota+2020
	Y21 = iota+2020
	Y22 = iota+2020
	Y23 = iota+2020
)

func main() {

	//Func scope variables declared using short declaration
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x,y,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(gx)
	fmt.Println(gy)
	fmt.Println(gz)

	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

	s := fmt.Sprintf("%d %s %t", x2, y2, z2)
	fmt.Println(s)


	//myType Var
	fmt.Println(myVar)
	fmt.Printf("%T\n",myVar)

	myVar = 42
	fmt.Println(myVar)

	myNewVar = int(myVar)
	fmt.Println("Converted Value",myNewVar)

	//Printing the decimal, binary and hex representation of an integer
	fmt.Printf("%d\t%b\t%#x\n", myNewVar, myNewVar, myNewVar)

	//Typed Constants
	fmt.Printf("Typed Constants %T %v\n",cx, cx)

	//UnType Constants
	fmt.Printf("UnTyped Constants %T %v\n",ucx, ucx)

	//Raw Literal String
	rs := `Hello   raw string
			"My name is James Bond"
			This is a multiline variable`
	nors := "Hello   raw string"
	fmt.Println(rs, nors)


	//IOTA values
	fmt.Println(Y20, Y21, Y22, Y23)
}
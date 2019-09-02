package main

import "fmt"

func main() {
	a := 10
	b := 20
	c:= 30

	d := a == 10
	e := b <= a
	f := b >= a
	g := b != c
	h := b < c
	i := b > c
	fmt.Println("a =",a,"b =",b,"c =",c)
	fmt.Println("\na == 10 is", d)
	fmt.Println("b <= a is",e)
	fmt.Println("b >= a is",f)
	fmt.Println("b != c is",g)
	fmt.Println("b < c is",h)
	fmt.Println("b > c is",i)

	bb := 4
	fmt.Printf("%d\t%b\t%#x\n", bb,bb,bb)
	bb1 := bb << 1
	fmt.Printf("%d\t%b\t%#x\n", bb1,bb1,bb1)
}
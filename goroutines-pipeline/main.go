package main

/* This file contains the functions created in golang to simulate a Pipeline functionality
for generating random numbers, sum it up and print them atlast.
*/

import (
	"fmt"
	"math/rand"
)

var CLOSEF bool

func first(n1 int, n2 int, f chan<- int) {
	for {
		if CLOSEF == true {
			close(f)
			return
		}
		randNum := rand.Intn(n2-n1+2)
		f<-randNum
	}
	
}

func second(s chan<- int, f <-chan int, n1 int, n2 int) {
	for val := range f {
		if val == (n2-n1) {
			CLOSEF = true
		}
		s<-val
	}
	close(s)
}

func third (s <-chan int) {
	for val := range s {
		fmt.Print(val, " ")
	}
	fmt.Println("\nDone Generating rand numbers")
}

func main() {
	var n1, n2 int
	f := make (chan int)
	s := make (chan int)
	CLOSEF = false

	fmt.Println("Input a range")
	fmt.Scanf("%d %d\n", &n1, &n2)

	go first(n1, n2, f)
	go second(s, f, n1, n2)
	third(s)
}

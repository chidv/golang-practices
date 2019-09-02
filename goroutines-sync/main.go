package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func first(n1 int, n2 int, f chan<- int) {
	defer wg.Done()
	for i:= n1; i< n2; i++ {
		f<-i
	}
	close(f)	
}

func second(f <-chan int) {
	
	defer wg.Done()
	sum := 0
	for val := range f {
		sum = sum + (val*val)
	}
	fmt.Println("Sum of all natural numbers in the given range is", sum)
}

func main() {
	var n1, n2 int
	f := make (chan int)

	fmt.Println("Input a range between which sum of natural numbers need to be generated")
	fmt.Scanf("%d %d\n", &n1, &n2)

	//What will happen if wg.Add is performed in the go routines
	wg.Add(1)
	go first(n1, n2, f)
	wg.Add(1)
	go second(f)
	wg.Wait()
}

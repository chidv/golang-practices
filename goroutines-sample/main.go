package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Starting with goroutines number:", runtime.NumGoroutine())

	for i:=1; i <=2; i++ {
		wg.Add(1)
		go func(x int) {
			for j:= x; j < 100; j+=2 {
				fmt.Println("Number Series :",j)
			}
			wg.Done()
		}(i)
	}
	fmt.Println("Before ending - goroutines number:", runtime.NumGoroutine())
	wg.Wait()
}
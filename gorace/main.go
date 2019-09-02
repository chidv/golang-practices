package main

//Build or Run with -race flag will show the race condition created by goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"runtime"
)

var counter int32

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i:=0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	fmt.Println("Number of goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter :", counter)


	counter = 0
	for i:=0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}
	fmt.Println("Number of goroutines", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter :", counter)

}
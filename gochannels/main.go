package main

import (
	"fmt"
)

func send(ch chan<- int) {
	for i:= 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	//Buffered Channels
	c := make(chan int, 1)

	c <- 42
	fmt.Println(<-c)

	//Another Channel
	ch := make(chan int)
	
	go send(ch)
	//Range receive channel
	for v := range ch {
		fmt.Println("Received ", v)
	}
	fmt.Println("Done")

	//Another Channel
	/*
	ch = make(chan int)
	
	go send(ch)
	//Range receive channel
	for {
		select {
		case v, ok := <- ch:
			if !ok {
				fmt.Println("Going to Exit")
				return
			}
			fmt.Println("Received", v)
		}
	}*/

	ch = make (chan int, 10)
	q := make (chan int, 10)
	//10 goroutines
	for i:=0; i < 10; i++ {
		go func(ii int) {
			for x := ii; x < 100; x+=10 {
				ch <- x
			}
			q<-i
		}(i)
	}

	var count int
	//Receive
	for {
		select {
		case v := <-ch:
			fmt.Println("Received",v)
		case <- q:
			count++
			if count == 10 {
				fmt.Println("Done receiving. Exiting..")
				return
			}
		}
	}
}
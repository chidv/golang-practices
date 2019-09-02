package main

import (
	"fmt"
	"sync"
	"bufio"
	"os"
	"strings"
	//"time"
)

const (
	MAX_WORKERS = 3
)

var lines = make (chan string, MAX_WORKERS)
var counts = make (chan int, MAX_WORKERS)

func worker(w *sync.WaitGroup) {
	defer w.Done()
	for line := range lines {
		//fmt.Println(line)
		counts <- len(strings.Split(line, " "))
	}
}

func spawnWorkers(w *sync.WaitGroup, count int, done chan<- interface{}) {
	for i:=0; i<count; i++ {
		w.Add(1)
		go worker(w)
	}
	w.Wait()
	done <- true
}

func main() {
	var w sync.WaitGroup
	if len(os.Args) != 2 {
		fmt.Println("Error: filename is not given")
		return
	}

	var done = make (chan interface{})
	
	go spawnWorkers(&w, MAX_WORKERS, done)
	
	go func (fileName string) {
		fp,err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error: File Open failed")
			return
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		close(lines)
	}(os.Args[1])

	sum := 0
	//val := 0
	for {
		select {
		case val := <- counts:
			sum = sum + val
		case <- done:
			fmt.Println(sum)
			close(counts)
			return
		}
	}
		
}
package main

import (
	"fmt"
	"time"
)

func worker(in <-chan int, out chan<- int) {
	for i := range in {
		fmt.Println("[worker]", i)
		out <- i * 2
		time.Sleep(time.Second)
	}
	out <- 0
}

func testWorker() {
	in := make(chan int)
	out := make(chan int, 3)
	go worker(in, out)
	in <- 1
	in <- 2
	in <- 3
	close(in)
	for o := range out {
		fmt.Println(o)
		if o == 0 {
			close(out)
		}
	}
}

func main() {
	testWorker()
}

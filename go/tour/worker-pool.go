package main

import (
	"fmt"
	"time"
)

func timeout() {
	ch := make(chan bool)
	go worker(w, jobs, results, ch)
	select{
	case res := <-ch:
		fmt.Println("done")
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func worker(id int, jobs <-chan int, results chan<- int, ch chan<- bool) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
		ch<-true
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w < 3; w++ {
		go timeout(w, jobs, results)
	}

	for j := 0; j < 9; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 0; r < 9; r++ {
		<-results
	}
}

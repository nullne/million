package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(1, 2, 3, 4)
	o1 := sq(in)
	o2 := sq(in)
	for o := range merge(o1, o2) {
		fmt.Println(o)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) (<-chan int){
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(in <-chan int){
		for i := range in {
			out <- i
		}
		wg.Done()
	}
	for _, c := range cs {
		wg.Add(1)
		go output(c)
	}
	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}

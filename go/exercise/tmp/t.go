package main

import (
	"fmt"
)

type res struct {
	output []byte
	code   int
}

type foo struct {
	results []res
}

func (f *foo) add(res res) {
	f.results = append(f.results, res)

}

func main() {
	chres := make(chan res)
	go func() {
		defer close(chres)
		chres <- res{[]byte("cool"), 1}
		chres <- res{[]byte("nice to meet you"), 1}
		chres <- res{[]byte("fuck"), 1}
	}()

	f := &foo{}
	for r := range chres {
		f.add(r)
	}
	for _, res := range f.results {
		fmt.Println(string(res.output))
	}

}

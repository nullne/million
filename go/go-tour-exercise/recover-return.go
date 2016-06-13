package main

import (
	"fmt"
	"errors"
)

func nice() (e error) {
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			e = errors.New("love you")
		}else{
		}
	}()
	ch := make(chan string, 1)
	ch <- "nie"
	tmp, ok := <-ch
	fmt.Println(tmp, ok)
	close(ch)
	tmp, ok = <-ch
	fmt.Println(tmp, ok)
	// ch <- "nice"
	return e
}

func main() {
	err := nice()
	fmt.Println(err)
}

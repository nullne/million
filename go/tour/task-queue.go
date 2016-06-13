package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("fuck oaaaaaaaaaaaaau")
	a := make(chan bool, 5)
	b := make(chan bool)
	go func() {
		var i int = 0
		for {
			a <- true
			go func(counter int) {
				if counter == 20 {
					b <- true
					return
				}
				go func() {
					<-time.After(time.Second * 5)
					<-a
					fmt.Println("timeout")
				}()
				r := rand.Intn(5) + 1
				time.Sleep(time.Second * time.Duration(r))
				fmt.Println("i:", counter, "	duration: ", r, time.Now())
				<-a
			}(i)

			//fmt.Println("-----------------", i, time.Now())
			i++
		}
	}()
	<-b
	/*
		for res := range a {
			fmt.Println(res, time.Now())
		}
	*/
	return
}

package main

import (
    "fmt"
)

func main() {
    c1, c2 := make(chan int), make(chan string)
    o := make(chan bool)

    go func() {
        for {
            select {
                case v, ok := <-c1:
                    if !ok {
                        fmt.Println("v1 stop")
                        o <- true
                        break
                    }
                    fmt.Println("v1: ",v)
                case v, ok := <-c2:
                    if !ok {
                        fmt.Println("v2 stop")
                        o <- true
                        break
                    }
                    fmt.Println("v2: ",v)
            }
        }
    }()

    c1 <- 1
    c2 <- "hi"
    c1 <- 3
    c2 <- "hello1"
    c2 <- "hello2"
    c2 <- "hello3"
    c2 <- "hello4"
    
    close(c1)
    //close(c2)

    <- o
}

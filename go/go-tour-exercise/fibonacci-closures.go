package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a := int(0)
    b := int(1)
    return func() int {
        tmp := b
        b = a + b
        a = tmp
        return b
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}



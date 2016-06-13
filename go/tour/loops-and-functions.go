package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := float64(1)
    var delta, ztmp float64 = z, z
    const p = 0.01  
    for delta >= p {
        z = z - ( z * z - x ) / ( 2 * z )
        delta = z - ztmp
        ztmp = z
        if delta < 0 {
            delta = - delta
        }
    }
    return z
}

func main() {
    fmt.Println(Sqrt(100000))
}



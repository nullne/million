package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        err := ErrNegativeSqrt(x)
        return 0, err
    }
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
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}



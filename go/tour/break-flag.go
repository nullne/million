// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import "fmt"

func main() {
	a := make([][]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			a[i][j] = i * j
		}
	}

OuterLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			switch a[i][j] {
			case 4:
				fmt.Println("4 found")
				break OuterLoop
			case 9:
				fmt.Println("9 found")
				break
			default:
				fmt.Println("not found")
			}
		}
	}
}

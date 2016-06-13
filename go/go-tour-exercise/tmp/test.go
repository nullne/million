package main

import (
	// "encoding/json"
	"fmt"
)

func foo(m map[int]int) {
	fmt.Printf("%p\n", m)
	m[1] = 5
}

type fuck struct {
	a *int
}

type opentsdbRequest struct {
	Start   int64
	End     int64
	Queries []struct {
		Aggregator string
		Metric     string
		Downsample string
		Filters    []struct {
			Type    string
			Tagk    string
			Filter  string
			GroupBy bool
		}
	}
}

func main() {
	a := []int{1, 2}
	a = append(a, nil...)
	fmt.Println(a)
	fmt.Printf("%s")
}

//numberSplit split large mix-up number into max and min for rate function
func numberSplit(val float64) (float64, float64) {
	tmin := int(val) % 10000
	tmax := float64(int(val)-tmin) / 10000
	return tmax / 100, float64(tmin) / 100
}

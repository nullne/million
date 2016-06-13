package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct{
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (re Rectangle) Area() float32 {
	return re.length * re.width
}

func main() {
	r := Rectangle{5,3}
	s := &Square{4}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	shapes := []Shaper{r, s}
	fmt.Println("looping through shapes for area...")
	for n, _ := range shapes{
		fmt.Println("Details", shapes[n])
		fmt.Println("Area:", shapes[n].Area())
	}
}

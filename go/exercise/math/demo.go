package main

import (
	"errors"
	"fmt"
)

const (
	M_AND = iota
	M_OR
)

//没有老考虑转义符号
func main() {
	// and start: 0, end: 102
	// or start: 35, end: 102
	// or start: 35, end: 102
	// str := ("( $1 > '2015-12-21 23:00:00' ) and (( 200 <= $2 < 500 ) or ( $3 ~= '(and)repattern') or ( $3 = 'fuck'))")

	// and 13 start: 0, end: 35
	// and 20 start: 0, end: 35
	// or 5 start: 0, end: 11
	// or 29 start: 24, end: 35
	str := "( $1 or ( $5 and $6 ) ) and $2 and ( $3 or $4 )"
	res, err := New(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res)

	and, or := positions(str)
	all := make(map[int][3]int, len(and)+len(or))

	for _, a := range and {
		start, end := border(str, a, 3)
		all[a] = [3]int{start, end, M_AND}
	}
	for _, o := range or {
		start, end := border(str, o, 2)
		all[o] = [3]int{start, end, M_OR}
	}
	fmt.Println(all)
	/*
			length := len(str)
			// var closed bool = true
			var QUOTES map[byte]bool = map[byte]bool{'\'': true, '"': true}
			// var PARENTHESES = [2]byte{'(', ')'}
			quotes := make(map[byte]bool, 2)
		Outer:
			for i, _ := range str {

				if QUOTES[str[i]] {
					if (i > 0 && str[i-1] != '\\') || i < 1 {
						//close start
						if quotes[str[i]] {
							quotes[str[i]] = false
						} else {
							quotes[str[i]] = true
						}
					}
				}
				for _, b := range quotes {
					if b {
						continue Outer
					}
				}

				if r := i + 3; r <= length && str[i:r] == "and" {
					fmt.Printf("and found: %d\n", i)
				}

				if r := i + 2; r <= length && str[i:r] == "or" {
					fmt.Printf("or found: %d\n", i)
				}
			}
	*/
}

type Expression struct {
	Left, Right string
	Operator    string
}

type Tree struct {
	Parent     *Tree
	ChildList []*Tree
	Value      interface{}
	T          int
}

func New(str string) (*Tree, error) {
	and, or := positions(str)
	helper := make(map[int][3]int, len(and)+len(or))
	for _, a := range and {
		start, end := border(str, a, 3)
		helper[a] = [3]int{start, end, M_AND}
	}
	for _, o := range or {
		start, end := border(str, o, 2)
		helper[o] = [3]int{start, end, M_OR}
	}


	length := len(helper)
	if length == 0 {
		return nil, errors.New("not found")
	}

	var root *Tree
	insert(root, str, helper)

	return root, nil
	var max int = -1
	roots := make(map[int][3]int, length)
	for h, border := range helper {
		r := border[2] - border[1]
		if r > max {
			for k, _ := range roots {
				delete(roots, k)
			}
			roots[h] = border
			max = r
			// remove
		} else if r == max {
			roots[h] = border
			// add
		}
	}

	var t int = -1
	for _, border := range roots {
		if t == -1 {
			t = border[2]
		} else if t != border[2] {
			return nil, errors.New("Illeagal Expression.")
		}
	}
	root.T = t
	root.Parent = parent

	return nil, nil
}


//helper function
func insert(t *Tree, str string, helper map[int][3]int) *Tree{
	if t == nil {
		return &Tree{}
	}

	for i, borders := range helper{
		child := &Tree{
			t,
			[]*Tree{},
			"",
			M_AND,
		}
		t.ChildList = append(t.ChildList, child)
	}
	max := maxHelper(helper)
	for _, m := range max{
		child := insert()
		t.ChildList = append(t.ChildList, child)
	}

	return nil
}
//positions returns position of and and or
func positions(str string) ([]int, []int) {
	var and, or []int
	length := len(str)
	QUOTES := map[byte]bool{'\'': true, '"': true}
	quotes := make(map[byte]bool, 2)
Outer:
	for i, _ := range str {
		if QUOTES[str[i]] {
			if (i > 0 && str[i-1] != '\\') || i < 1 {
				//close start
				if quotes[str[i]] {
					quotes[str[i]] = false
				} else {
					quotes[str[i]] = true
				}
			}
		}
		for _, b := range quotes {
			if b {
				continue Outer
			}
		}

		if r := i + 3; r <= length && str[i:r] == "and" {
			and = append(and, i)
		} else if r := i + 2; r <= length && str[i:r] == "or" {
			or = append(or, i)
		}
	}
	return and, or
}

// border returns start and end for each index
func border(str string, index, width int) (int, int) {
	length := len(str)
	start := 0
	end := length - 1
	counter := 0
	for i := index - 1; i >= 0; i-- {
		if str[i] == ')' {
			counter++
		}
		if str[i] == '(' {
			counter--
		}
		if counter < 0 {
			start = i
			break
		}
	}

	counter = 0
	for i := index + width; i < length; i++ {
		if str[i] == '(' {
			counter++
		}
		if str[i] == ')' {
			counter--
		}
		if counter < 0 {
			end = i
			break
		}
	}
	return start, end
}

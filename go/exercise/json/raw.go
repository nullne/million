package main

import (
	"encoding/json"
	"fmt"
)

type foo struct {
	I int
	F interface{}
}

type fuck struct {
	I string
}

func main() {
	f := foo{1, map[string]int{"fuck": 1, "you": 2}}
	// f := fuck{"fuck"}
	b, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Printf("|%s|\n", b)

	var ss foo
	fmt.Println(json.Unmarshal(b, &ss))
	fmt.Println(ss)
}

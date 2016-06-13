package main

import (
	"encoding/json"
	// "errors"
	"fmt"
)

type foo struct {
	I   *int
	M   string
	Err MyError
}
type MyError struct {
	Code    int
	Message string
}

func (err MyError) Error() string {
	return "fuck"
}

type Foo struct {
	I int
	M string
}

func (f Foo)Decode(body []byte)error{
	return json.Unmarshal(body, f)
}

func test() error {
	return MyError{1, "fuck"}
}
func main() {
	m := Foo{1, "fuck"}
	body, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
	// x := Foo{}
	// err = x.Decode([]byte(body))
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(x)
}

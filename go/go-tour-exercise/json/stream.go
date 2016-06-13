package main

import (
	"encoding/json"
	// "errors"
	"fmt"
	"bufio"
	"bytes"
)
type Foo struct {
	I int
	M string
}

func main(){
	f := Foo{I: 1}
	b := bytes.Buffer{}
	foo := bufio.NewWriter(&b)
	encoder := json.NewEncoder(foo)
	err := encoder.Encode(f)
	if err != nil {
		panic(err)
	}
	err = foo.Flush()
	if err != nil {
		panic(err)
	}
	var x Foo
	json.Unmarshal(b.Bytes(), &x)
	fmt.Println(x)

}

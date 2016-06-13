package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type foo struct {
	A os.FileMode    `json:"a"`
	S string `json:"s"`
}

func main() {
	// f := foo{0644, "fuck"}
    //
	// b, err := json.Marshal(&f)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(b))
	
	b := `{"a":644,"s":"fuck"}`
	var n foo
	if err := json.Unmarshal([]byte(b), &n); err != nil {
		panic(err)
	}
	fmt.Println(n)

}

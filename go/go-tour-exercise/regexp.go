package main

import (
	"fmt"
	"regexp"
	"strconv"
	"os"
)

func handle(origin []byte, re *regexp.Regexp)(rtn []byte,code int){
	res := re.FindAllSubmatch(origin, -1)
	if len(res) == 0 {
		return origin, -1
	}
	code, err := strconv.Atoi(string(res[0][1]))
	if err != nil {
		panic(err)
	}
	rtn = re.ReplaceAllLiteral(origin, []byte{})
	return
}

func main() {
	input := []byte("/root\n\r<|123|>")
	re := regexp.MustCompile(`<\|(\d+)\|>`)
	fmt.Println(handle(input, re))
	os.Exit(-1)
}


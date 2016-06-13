package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd := flag.String("c", "", "cmd")
	flag.Parse()
	for _, l := range *cmd {
		fmt.Println(l)
	}
}

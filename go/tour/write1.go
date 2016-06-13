package main

import (
	"fmt"
	"io"
	"os"
)

func read1(file string) string {
	fi, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	chuncks := make([]byte, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		chuncks = append(chuncks, buf...)
	}
	return string(chuncks)
}

func main() {
	fmt.Println(read1("./a"))
}

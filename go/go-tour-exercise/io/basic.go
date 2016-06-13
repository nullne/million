package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("/tmp/dat")
	check(err)

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n", n3, o3, string(b3))

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	b5 := make([]byte, 10)
	n5, err := f.Read(b5)
	fmt.Printf("%d bytes @ fuck: %s \n", n5, string(b5))

	b5 = make([]byte, 10)
	n5, err = f.Read(b5)
	fmt.Printf("%d bytes @ fuck: %s \n", n5, string(b5))

	f.Close()

}

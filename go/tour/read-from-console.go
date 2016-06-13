package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	menu := `
1. hi
2. fuck
3. quit
Please select:`
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(menu)
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n\r ")
		switch text {
		case "1":
			fmt.Printf("123")
		case "2":
			fmt.Printf("223")
		case "q":
			return
		}
		fmt.Printf("%q", text)
	}
}

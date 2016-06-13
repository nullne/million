package main

import (
	"bufio"
	// "errors"
	"fmt"
	// "io/ioutil"
	"os"
	"time"
	// "strings"
	// "unicode"
	// "unicode/utf8"
)

func main(){
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())
	file, err := os.Open("/Users/nullne/Desktop/iTunes_IP_WHITELIST_logProcessor.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text(), "_____")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

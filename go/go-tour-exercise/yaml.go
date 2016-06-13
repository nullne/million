package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Nums struct {
	One   string
	Two   string
	Three string
	Four  string
}
type Config struct {
	Foo string
	Bar []Nums
}

func main() {
	var config Config
	source, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Print(config)
	// fmt.Printf("Value: %#v\n", config.Bar)
}

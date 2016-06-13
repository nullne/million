package main

import(
	"fmt"
)

type Animal interface{
	Speak() string
}

type Cat struct{
}

func (c *Cat) Speak() string{
	return "miao"
}

type Dog struct {
}

func (d Dog) Speak() string{
	return "wang"
}

func printAll(vals []interface{}){
	for _, val := range vals{
		fmt.Println(val)
	}
}
func main() {
	names := []string{"nice", "to", "meet", "you"}
	vals := make([]interface{}, len(names))
	for i, v := range names{
		vals[i] = v
	}
	printAll(vals)
	fmt.Println("-----------------")
	animals := []Animal{new(Cat), Dog{}}
	// animals := []Animal{&Cat{}, Dog{}}
	for _, animal := range animals{
		fmt.Println(animal.Speak())
	}
}

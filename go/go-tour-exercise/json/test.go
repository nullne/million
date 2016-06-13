package main

import(
	"os"
)

func main(){
	f, err := os.Open("./foo")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := f.Chmod(888888); err != nil {
		panic(err)
	}

}

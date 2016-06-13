package main

import (
	"text/template"
	"io"
	"io/ioutil"
	"fmt"
)

type Person struct {
	// Name string //exported field since it begins with a capital letter
	WorkingPath, BinUrl, BinName, BinMd5Sum string
}

func main() {
	r, w := io.Pipe()
	// t := template.New("emplate") //create a new template with some name
	// t, _ = t.Parse("hello {{.Name}}!") //parse some content and generate a template, which is an internal representation
	t, _ := template.ParseFiles("./log.template")
	// .ParseFiles("./log.template")
	p := Person{"/work/path", "www.baidu.com", "binname", "md5sum"} //define an instance with required field
	go func(){
		t.Execute(w, p)                              //merge template ‘t’ with content of ‘p’
		w.Close()
	}()
	fmt.Println(ioutil.ReadAll(r))
}

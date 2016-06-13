package main

import (
	"bytes"
	"text/template"
	"fmt"
)

type Machine struct {
	Host string
}

func main() {
	var b bytes.Buffer
	machine := &Machine{"fuck"}
	tpl := "{{ .Host }}"
	t := template.Must(template.New("t").Parse(tpl))
	err := t.Execute(&b, machine)
	if err != nil {
		panic(err)
	}
	fmt.Println(b.String())
}

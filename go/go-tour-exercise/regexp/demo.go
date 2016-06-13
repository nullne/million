package main

import(
	"fmt"
	"regexp"
	// "bufio"
	"bytes"
)

func main(){
	a := bytes.NewBuffer([]byte(""))
	a.String()
	//log naming
	// re := regexp.MustCompile(`(\w*?)-access\.log\.(\w*?)\.(.*?)\.([a-zA-Z]{3}-[a-zA-Z]{2}-[\w]{1}-[\w]{3})\.(\d*?)\.gz`)
	// match := re.FindStringSubmatch("fc-access.log.baidu.07002713g8.CMN-WH-1-3g8.20151209185156.gz")
	// for i, val := range match{
	// 	fmt.Printf("%v: %v\n", i, val)
	// }

	re := regexp.MustCompile(`<\|(\d+)-(\d*)\|>`)
	ss := re.FindAllSubmatch([]byte("<|222-p22|><|333-p|>"), -1)
	fmt.Println(ss == nil)
	for i, s := range ss {
		fmt.Printf("\n%d: %s", i, s)
	}
}

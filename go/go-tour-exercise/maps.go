package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    rtn := make(map[string]int)
    for _, value := range strings.Fields(s){
        _, ok := rtn[value]
        if ok {
            rtn[value] ++
        } else {
            rtn[value] = int(1)
        }
    }
    return rtn
}

func main() {
    wc.Test(WordCount)
}



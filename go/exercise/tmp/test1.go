package main

import (
	"fmt"
	"regexp"
	"strings"
)

type foo struct {
	parent *foo
	value  int
}

func test(t []int) {
	fmt.Printf("%p: %v\n", &t, &t)
	t = append(t, 1)
	fmt.Printf("%p: %v\n", t, t)
}

type Spliter interface {
	Split(ori string) []string
}

type RegSpliter struct {
	Pattern string
	pattern *regexp.Regexp
	N       int
}

func NewRegSpliter(p string, n int) *RegSpliter {
	var err error
	s := &RegSpliter{Pattern: p, N: n}
	s.pattern, err = regexp.Compile(s.Pattern)
	if err != nil {
		return nil
	}
	return s
}

func (r RegSpliter) Split(ori string) []string {
	ss := r.pattern.FindStringSubmatch(ori)
	if len(ss) == (r.N + 1) {
		return ss[1:]
	}
	return nil
}

type SpaceSpliter struct {
	N int
}

// @TODO num of space after n may be diminished
func (s SpaceSpliter) Split(ori string) []string {
	ss := strings.Fields(ori)
	l := len(ss)
	switch {
	case l < s.N:
		return nil
	case l == s.N:
		return ss
	case l > s.N:
		last := strings.Join(ss[s.N:], " ")
		return append(ss[:s.N], last)
	}
	return nil
}

func main() {
	m := RegSpliter{}
	fmt.Println(m == RegSpliter{})
	// test := "-axxxbyc-"
	// test := "-abzc-"
	// test := "-abzc-"
	test := "fuck you"
	s := NewRegSpliter("a(x*)b(y|z)c", 2)
	if s == nil {
		fmt.Println("Nothing")
	}
	fmt.Println(s.Split(test))
}

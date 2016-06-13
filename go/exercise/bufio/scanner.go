package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Tree struct {
	Left, Right *Tree
	//0 => node, operand works
	//1 => leave, expression works
	Type       uint
	Operand    string
	Expression *Expression
}
type Expression struct {
	Operator string
	Operand  string
	operand  interface{}
	// T        ColumnType
}

//supported operator
const (
	O_EQUAL         = "="
	O_GREATER       = ">"
	O_GREATER_EQUAL = ">="
	O_LESS          = "<"
	O_LESS_EQUAL    = "<="
	O_REG_EQUAL     = "~="
	O_NOT_EQUAL     = "!="
)

//logical operator
const (
	LO_AND = "and"
	LO_OR  = "or"
	LO_XOR = "xor"
	LO_NOT = "not"
)

func logicalOpts() []string {
	return []string{
		LO_AND,
		LO_OR,
		LO_XOR,
		LO_NOT,
	}
}

func arithmeticOpts() []string {
	return []string{
		O_EQUAL,
		O_GREATER,
		O_GREATER_EQUAL,
		O_LESS,
		O_LESS_EQUAL,
		O_REG_EQUAL,
		O_NOT_EQUAL,
	}
}

func main() {
	condition := `(($0>100) and (not $1<=200)) or ($3~= "fuck" xor $5!=300)`
	fmt.Println(condition)
	s := bufio.NewScanner(strings.NewReader(condition))
	s.Split(splitFunc)
	var postfix []string
	var stack []string
	var opd []string
	var opt string
Outer:
	for s.Scan() {
		fmt.Printf("token: |%s|\n", s.Text())
		token := s.Text()
		switch {
		case token == " ":
			continue Outer
		case isOpt(token, false):
			opt = token
		case isOpt(token, true) || token == "(" || token == ")":
			for len(stack) != 0 && opOrder(stack[len(stack)-1], token) {
				var op string
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				postfix = append(postfix, op)
			}
			if len(stack) == 0 || token != ")" {
				stack = append(stack, token)
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
			opd = append(opd, token)
			if len(opd) == 2 {
				//new expression
				if opt != "" {
					postfix = append(postfix, fmt.Sprintf("%s %s %s", opd[0], opd[1], opt))
				} else {
					fmt.Println("F**************k")
					fmt.Println(opd)
				}
				//
				opt = ""
				opd = opd[:0]
			}
		}
	}
	if len(stack) != 0 {
		postfix = append(postfix, stack...)
	}
	for _, tok := range postfix {
		fmt.Println(tok)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
}

func isOpt(s string, logical bool) bool {
	var opts []string
	if logical {
		opts = logicalOpts()
	} else {
		opts = arithmeticOpts()
	}
	for _, o := range opts {
		if o == s {
			return true
		}
	}
	return false
}

func opOrder(op1, op2 string) bool {
	order := map[string]int{
		"and": 1,
		"or":  1,
		"xor": 1,
		"not": 2,
	}
	if op1 == "(" || op2 == "(" {
		return false
	} else if op2 == ")" {
		return true
	} else {
		if order[op1] < order[op2] {
			return false
		} else {
			return true
		}
	}
}

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	switch data[0] {
	case ')', '(':
		advance, token, err = 1, data[:1], nil
	case '"', '\'':
		advance, token, err = consumeString(data, atEOF)
	// case ' ', '\t':
	// 	advance, token, err = 1, data[:1], nil
	default:
		advance, token, err = consumeWord(data, atEOF)
	}
	return
}

func consumeString(data []byte, atEOF bool) (int, []byte, error) {
	delim := data[0]
	skip := false
	end := false
	accum := []byte{}
	for i, b := range data[1:] {
		if b == delim && !skip {
			return i + 2, accum, nil
		}
		skip = false
		if b == '\\' {
			skip = true
			continue
		}
		accum = append(accum, b)
		/*
		   if i == (len(data) - 1) {
		       return i + 1, accum, nil
		   }
		*/
	}
	if atEOF && !end {
		return len(data), accum, errors.New("fuck")
	}
	return 0, nil, nil
}

//and not xor
//or >= <= ~=
//> < =
func endWithOperator(data []byte, n int) int {
	if n > 1 {
		if string(data[n-2:n+1]) == "and" || string(data[n-2:n+1]) == "not" || string(data[n-2:n+1]) == "xor" {
			return 2
		}
	}
	if n > 0 {
		if string(data[n-1:n+1]) == "or" || string(data[n-1:n+1]) == ">=" || string(data[n-1:n+1]) == "<=" || string(data[n-1:n+1]) == "~=" || string(data[n-1:n+1]) == "!=" {
			return 1
		}
	}
	if data[n] == '>' || data[n] == '=' || data[n] == '<' || data[n] == '(' || data[n] == ')' {
		return 0
	}
	return -1
}

func startWithOperator(data []byte) int {
	if string(data[0:3]) == "and" || string(data[0:3]) == "not" || string(data[0:3]) == "xor" {
		return 3
	}
	if string(data[0:2]) == "or" || string(data[0:2]) == ">=" || string(data[0:2]) == "<=" || string(data[0:2]) == "~=" || string(data[0:2]) == "!=" {
		return 2
	}
	if data[0] == '>' || data[0] == '=' || data[0] == '<' || data[0] == '(' || data[0] == ')' {
		return 1
	}
	return 0
}

func consumeWord(data []byte, atEOF bool) (int, []byte, error) {
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}
	if start != 0 {
		return start, []byte{' '}, nil
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) {
			return i + width, data[start:i], nil
		}
		if back := endWithOperator(data[start:], i-start); back >= 0 {
			if advance := startWithOperator(data[start:]); advance > 0 {
				return start + advance, data[start : start+advance], nil
			} else {
				return i - back, data[start : i-back], nil
			}
		}
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

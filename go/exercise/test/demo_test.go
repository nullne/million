package demo

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("division function tests do not pass")
	}else {
		t.Log("first test passed")
	}
}

func Test_Divisiono_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		t.Error("Division did not work as expected.")
	}else {
		t.Log("one test passed", e)
	}
}

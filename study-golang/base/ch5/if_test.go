package ch5

import "testing"

func someFun() (int, error) {
	return 1, nil

}

func TestIf(t *testing.T) {
	if v, err := someFun(); err == nil {
		t.Log(v)
	} else {
		t.Log(err)
	}
}

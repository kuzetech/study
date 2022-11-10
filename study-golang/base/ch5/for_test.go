package ch5

import "testing"

func TestFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(i)
	}
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestWhileTrue(t *testing.T) {
	for {
		t.Log(1)
	}
}

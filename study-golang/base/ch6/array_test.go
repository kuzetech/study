package ch6

import "testing"

func TestArrayDefine(t *testing.T) {
	var a [3]int
	a[0] = 1
	t.Log(a)

	b := [3]int{1, 2, 3}
	t.Log(b)

	c := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	t.Log(c)

	d := [...]int{1, 2, 3}
	t.Log(d)
}

func TestArrayFor(t *testing.T) {
	var a []int = []int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		t.Log(a[i])
	}

	for index, value := range a {
		t.Log(index, value)
	}

	for _, value := range a {
		t.Log(value)
	}
}

func TestArraySplit(t *testing.T) {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	t.Log(a[1:2])
	t.Log(a[1:len(a)])
	t.Log(a[1:])
	t.Log(a[:2])
}

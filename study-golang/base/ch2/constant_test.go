package ch2

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant1(t *testing.T) {
	t.Log(Monday)
	t.Log(Tuesday)
	t.Log(Wednesday)
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Executable)
}

func TestConstant2(t *testing.T) {
	a := 7 // 111
	t.Log(a & Readable)
	t.Log(a & Writable)
	t.Log(a & Executable)

}

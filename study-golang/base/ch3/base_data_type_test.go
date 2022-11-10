package ch3

import (
	"math"
	"testing"
)

/*
特别注意：
	不允许隐式类型转换，即使是别名也不允许

基本类型：
	bool
	string
	int(看机器是32还是64)  int8 int16 int32 int64
	uint(看机器是32还是64) uint8 uint16 uint32 uint64 uintptr
	byte = alias for uint8
	rune = alias for int32, represents a unicode code point
	float32 float64
	complex64 complex128
*/

type MyInt int64 // 定义一个别名

func TestImplicitConversion(t *testing.T) {
	//var a MyInt = 1
	//var b int64 = 2
	//a = b
	//
	//var c int = 1
	//var d int64 = 2
	//c = d
}

func TestMaxValue(t *testing.T) {
	t.Log(math.MaxInt64)
	t.Log(math.MaxFloat64)
}

// 指针类型
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Log("%T %T", a, aPtr)
	// aPtr = aPtr + 1	// 不支持指针的任何运算
}

func TestString(t *testing.T) {
	var s string
	if s == "" {
		t.Log("字符串的初始化是空字符串")
	} else {
		t.Log("字符串的初始化不是空字符串")
	}
}

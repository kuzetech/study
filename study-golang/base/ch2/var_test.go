package ch2

import (
	"fmt"
	"testing"
)

func TestFirstTry(t *testing.T) {
	var a int = 1
	b := 2 // 类型推断
	var (
		c int = 1
		d     = 1 //类型推断
	)

	a, b = b, a // 两个变量交换值

	t.Log(a)
	t.Log(b)
	t.Log(c)
	t.Log(d)

}

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	var temp = 0
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		temp = b
		b = a + b
		a = temp
	}
}

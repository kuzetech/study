package ch6

import "testing"

func TestSliceDefine(t *testing.T) {

	// 第一种定义方式
	var s0 []int // 初始化的时候，和数组的区别仅在于不指定长度
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	// 第二种定义方式
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// 第三种定义方式
	s2 := make([]int, 3, 5) // length = 3 ， cap = 5，仅初始化前三个元素
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
}

func TestSliceGrowing(t *testing.T) {
	// 当存储空间不够了， cap * 2
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	Q2 := year[3:6]
	t.Log(len(Q2), cap(Q2)) // 3,9 cap 是 year 从3开始的长度
	summer := year[5:8]
	t.Log(len(summer), cap(summer)) // 3,7

	summer[0] = 111 // 因为共享了切片，一个修改，大家都变
	t.Log(Q2)
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	// 切片只能跟 nil 进行比较
}

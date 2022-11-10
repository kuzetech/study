package ch4

import "testing"

// 其他语言的数组比较，比较的是引用的地址
// 在 go 中，数组的维数相同，元素个数相同才可以比较，否则报错
// 并且每一个元素都相同的才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// t.Log(a == c) // 编译报错
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7        // 111
	b := a &^ 110 // 对应位上，如果0保留，如果1清零,第一位不能是零否则无效
	t.Log(b)      // 0101
}

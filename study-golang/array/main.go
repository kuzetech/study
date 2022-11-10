package array

import "log"

func test() {
	// 不初始化
	array := make([]int, 0, 5)
	log.Println(array)
	array = append(array, 1)
	log.Println(array)
}

type S struct {
	a int64
	b int64
}

func Test2() {
	// 不初始化
	array := make([]S, 5)
	log.Println(array)
	array = append(array, S{a: 1})
	log.Println(array)
}

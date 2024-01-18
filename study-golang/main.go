package main

import "fmt"

func main() {
	// 原始数组
	originalArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	// 将数组四等分
	length := len(originalArray)
	quarterSize := length / 4

	// 使用 for 循环切片
	for i := 0; i < 4; i++ {
		start := i * quarterSize
		end := (i + 1) * quarterSize

		// 防止索引越界
		if end > length {
			end = length
		}

		// 切片
		slice := originalArray[start:end]

		// 打印结果
		fmt.Printf("Slice %d: %v\n", i+1, slice)
	}
}

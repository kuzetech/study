package complementZero

import "fmt"

func test() {
	var i = 120
	newStr := fmt.Sprintf("%05d", i)
	fmt.Println(newStr) // 0001
}

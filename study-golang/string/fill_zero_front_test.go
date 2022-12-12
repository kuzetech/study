package string

import (
	"fmt"
	"testing"
)

func Test_fill_zero_front(t *testing.T) {
	var i = 120
	newStr := fmt.Sprintf("%05d", i)
	fmt.Println(newStr) // 00120
}

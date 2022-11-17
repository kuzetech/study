package flag

import (
	"fmt"
	"os"
)

func test() {
	args := os.Args
	for _, arg := range args {
		fmt.Printf("系统传递参数为 %s", arg)
		println()
	}
}

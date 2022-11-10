package ch5

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitchString(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin", "test":
		fmt.Println("win")
		// 自带 break
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("other")
	}
}

func TestSwitchIf(t *testing.T) {
	var num int = 0
	switch {
	case 0 <= num && num <= 3:
		fmt.Println("1-3")
	case 4 <= num && num <= 6:
		fmt.Println("4-6")
	default:
		fmt.Println("other")
	}
}

package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type Test struct {
	Items []*int
}

var t1 int = 3
var t2 int = 1
var t3 int = 2

var test = Test{
	Items: []*int{&t1, &t2, &t3},
}

func TestArrayFieldOrder(t *testing.T) {
	data, _ := jsoniter.ConfigFastest.Marshal(test)

	for i := 0; i < 100; i++ {
		var test2 = Test{}

		jsoniter.ConfigFastest.Unmarshal(data, &test2)

		for _, item := range test2.Items {
			fmt.Print(*item)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

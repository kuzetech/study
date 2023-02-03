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
	// json 序列化必须保证 数组的顺序
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

type Message struct {
	Time  int64 `name:"time" json:"time"`
	Count int64 `name:"count" json:"count"`
}

func TestBigIntDecode(t *testing.T) {

	content := "{\"time\": 1675404945109,\"count\": 1621071533109678080}"

	var result = Message{}
	jsoniter.ConfigFastest.Unmarshal([]byte(content), &result)

	fmt.Println(result)

}

func TestBigIntEncode(t *testing.T) {
	var msg = Message{
		Time:  1675404945109,
		Count: 1621071533109678080,
	}

	data, err := jsoniter.ConfigFastest.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}

	var result = Message{}
	jsoniter.ConfigFastest.Unmarshal(data, &result)

	fmt.Println(result)
}

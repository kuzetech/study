package json

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"testing"
)

func Test_value_type(t *testing.T) {
	fastApi := jsoniter.ConfigFastest

	bytes, err := ioutil.ReadFile("./complex.json")
	if err != nil {
		t.Fatal(err)
	}

	// 返回的类型 any 很方便，可以遍历 key；也可以随意转换 类型
	result1 := fastApi.Get(bytes, "person")
	result2 := fastApi.Get(bytes, "person", "name")
	result3 := fastApi.Get(bytes, "person", "name", "first")

	result4 := fastApi.Get(bytes, "person", "avatars")
	result5 := fastApi.Get(bytes, "person", "avatars", 0)
	result6 := fastApi.Get(bytes, "person", "avatars", 0, "url")

	t.Log(result1, result2, result3, result4, result5, result6)

}

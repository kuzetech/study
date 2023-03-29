package jsoniter

import (
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"testing"
)

func Test_unmarshal_bytes(t *testing.T) {
	// The default performance is already several times faster than the standard library
	// 如果直接使用 jsoniter.Marshal() 就使用的这一配置
	// defaultApi := jsoniter.ConfigDefault

	// tries to be 100% compatible with standard library behavior
	// compatibleApi := jsoniter.ConfigCompatibleWithStandardLibrary

	// If you want to have absolutely best performance
	// this will marshal the float with 6 digits precision (lossy), which is significantly faster
	fastApi := jsoniter.ConfigFastest

	bytes, err := ioutil.ReadFile("./data.json")
	if err != nil {
		t.Fatal(err)
	}

	person := &Person{}
	err = fastApi.Unmarshal(bytes, person)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(person.Last)

}

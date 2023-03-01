package json

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

type Message struct {
	Name  string  `name:"name" json:"name"`
	Count float64 `name:"count" json:"count"`
}

var contentBytes = []byte("{\"name\": \"test\",\"count\": 1621071533109678080123456}")

func TestBigIntDecodeUseFloat(t *testing.T) {
	var result interface{}
	// ConfigFastest marshals float with only 6 digits precision
	// will lose precession
	FloatEncoding.Unmarshal(contentBytes, &result)

	// map[count:1.621071533109678e+18 name:test]
	// 这里 count 字段明显是 float64 类型，已经丢失精度
	fmt.Println(result)
}

func TestBigIntDecodeUseNumber(t *testing.T) {
	var result interface{}
	// readNumberAsString
	NumberEncoding.Unmarshal(contentBytes, &result)

	// map[count:1621071533109678080123456 name:test]
	// 这里 count 字段实际类型为 json.Number，是 string 的别名
	fmt.Println(result)

	m := result.(map[string]interface{})
	count := m["count"]

	switch count.(type) {
	case float64:
		println("float64")
	case string:
		println("string")
	case json.Number:
		println("json.Number")
	default:
		println("unknown")
	}
}

func TestBigIntDecodeUseNumberToObject(t *testing.T) {
	var msg Message
	// readNumberAsString
	err := NumberEncoding.Unmarshal(contentBytes, &msg)

	// 如果我直接反序列化到对象中，count 字段将超过 int64 的接收范围，这时候将报错
	if err != nil {
		// json.Message.Count: readUint64: overflow, error found in #10 byte of
		// ...|: 162107153310967808|..., bigger context ...|{"name": "test","count": 1621071533109678080123456}|...
		t.Log(err)
	} else {
		log.Println(msg)
	}
}

func TestBigIntDecodeConvertUseNumberToObject(t *testing.T) {
	should := require.New(t)

	// 尝试将 json.Number 转换成 int64 范围
	// 如果出错将转换成 float64 ，会丢失精度
	data, err := DecodeJsonBytes(contentBytes)
	should.Nil(err)

	// map[count:1.621071533109678e+24 name:test]
	log.Println(data)

	m := data.(map[string]interface{})
	name := m["name"].(string)
	// 最终结果 count = float64 1621071533109678000000000
	count := m["count"].(float64)

	// 再将 interface 放入到 object 中
	var msg = Message{
		Name: name,
		// 当 float64 的值大于 int64 的值时会造成内存溢出
		// int64 最大值为 9223372036854775808
		Count: count,
	}

	// 最终结果 {test -9223372036854775808}
	log.Println(msg)
}

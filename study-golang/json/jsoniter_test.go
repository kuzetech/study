package json

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
	"log"
	"strconv"
	"strings"
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

func TestBigIntDecode(t *testing.T) {

	content := "{\"name\": \"test\",\"count\": 1621071533109678080}"

	var result interface{}
	// ConfigFastest marshals float with only 6 digits precision
	// will lose precession
	jsoniter.ConfigFastest.Unmarshal([]byte(content), &result)

	// map[count:1.621071533109678e+18 name:test]
	// 这里 count 字段明显是 float64 类型，已经丢失精度
	fmt.Println(result)

}

var JsonEncoding = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()

type Message struct {
	Name  string `name:"name" json:"name"`
	Count int64  `name:"count" json:"count"`
}

func TestBigIntDecodeUseNumber(t *testing.T) {
	should := require.New(t)

	content := "{\"#op\": 1}"

	var result interface{}
	// readNumberAsString
	JsonEncoding.Unmarshal([]byte(content), &result)

	// map[count:1621071533109678080123456 name:test]
	// 这里 count 字段实际类型为 json.Number，是 string 的别名
	fmt.Println(result)

	m, ok := result.(map[string]interface{})
	should.Equal(true, ok)

	count, exist := m["count"]
	should.Equal(true, exist)

	switch count.(type) {
	case float64:
		println("float64")
	case string:
		println("string")
	case jsonNumber:
		println("jsonNumber")
	case json.Number:
		println("json.Number")
	default:
		println("unknown")
	}

	// 如果我直接反序列化到对象中，count 字段将超过 int64 的接收范围
	var msg Message
	err := JsonEncoding.Unmarshal([]byte(content), &msg)
	// json.Message.Count: readUint64: overflow, error found in #10
	should.NotNil(err)

	// 由于 count 字段超过 int64 范围，将强制转换成 float64 丢失精度
	data, err := DecodeJsonBytes([]byte(content))
	should.Nil(err)

	m2, ok := data.(map[string]interface{})
	should.Equal(true, ok)

	// float64 1621071533109678000000000

	msg = Message{
		Name:  m2["name"].(string),
		Count: int64(m2["count"].(float64)),
	}

	log.Println(msg)

}

func DecodeJsonBytes(b []byte) (interface{}, error) {
	var decoded interface{}
	err := JsonEncoding.Unmarshal(b, &decoded)
	if err != nil {
		return nil, err
	}

	return convertNumbers(decoded), nil
}

func DecodeJsonString(s string) (interface{}, error) {
	return DecodeJsonBytes([]byte(s))
}

type jsonNumber interface {
	Float64() (float64, error)
	Int64() (int64, error)
	String() string
}

func convertNumbers(v interface{}) interface{} {
	switch vv := v.(type) {
	case []interface{}:
		for index, mp := range vv {
			vv[index] = convertNumbers(mp)
		}
	case map[string]interface{}:
		for k, v := range vv {
			vv[k] = convertNumbers(v)
		}
	default:
		v = convertDefault(v)
	}
	return v
}

func convertDefault(v interface{}) interface{} {
	switch vv := v.(type) {
	case jsonNumber:
		i, err := vv.Int64()
		if err == nil {
			return i
		}

		// judge suffix '.0' then try to parse prefix as a int64
		ok := strings.HasSuffix(vv.String(), ".0")
		if ok {
			i2, err := strconv.ParseInt(vv.String()[:len(vv.String())-2], 10, 64)
			if err != nil {
				panic("try parse prefix as a int64, but failed")
			}
			return i2
		}

		f, err := vv.Float64()
		if err != nil {
			panic("json.Number contain neither integer nor float")
		}
		return f
	default:
		return v
	}
}

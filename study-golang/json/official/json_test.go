package official

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Movie struct {
	title string `json:"title"` // 自动忽略小写字母开头的字断
	Year  int
	F1    bool   `json:"F1,omitempty"` // 加上 omitempty 如果类型是布尔值，值为 false 将被忽略；如果是其他类型，为 nil 将被忽略
	F2    bool   `json:"F2"`
	F3    string `json:"-"`  // 表示忽略该字段
	F4    string `json:"-,"` // 表示 json 字段名为 -
}

func Test_json(t *testing.T) {
	movie := Movie{
		title: "Casablanca",
		Year:  1942,
		F1:    false,
		F2:    false,
		F3:    "Casablanca",
		F4:    "Casablanca",
	}

	// 序列化
	data, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("序列化的结果为：%s \n", data)

	var d Movie
	err = json.Unmarshal(data, &d)

	if err != nil {
		log.Fatalf("JSON Unmarshal failed: %s", err)
	}

	fmt.Println("反序列化的结果如下：")
	log.Println(d)

}

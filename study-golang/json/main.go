package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	title string `json:"title"`
	Year  int
	F1    bool   `json:"F1,omitempty"` // 加上 omitempty 如果值为 false 将被忽略，如果是其他类型，为 nil 也将被忽略
	F2    bool   `json:"F2"`
	F3    string `json:"-"` // 如果内容只包含 - ，则表示忽略该字段，需要特别注意，如果内容为 -, 则表示 json 字段名为 -
}

func test() {
	movie := Movie{
		title: "Casablanca",
		Year:  1942,
		F1:    false,
		F2:    false,
		F3:    "Casablanca",
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

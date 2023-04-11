package file

import (
	"bufio"
	"io"
	"log"
	"os"
	"testing"
)

func init() {
	file, _ := os.Create("./files/read.txt")
	file.Write([]byte("\n"))
	file.Write([]byte("1111\n"))
	file.Write([]byte("2222\n"))
	file.Write([]byte("3333"))
	file.Close()
}

func Test_read(t *testing.T) {
	file, _ := os.Open("./files/read.txt")
	r := bufio.NewReader(file)
	for {
		lineBytes, terr := r.ReadBytes('\n')
		// 第三次同时返回 3333 和 eof，因此需要选处理数据
		if len(lineBytes) > 0 {
			// 删除输出 '\n'
			// log.Println("读取到的内容为：" + string(lineBytes[:(len(lineBytes)-1)]))
			log.Println("读取到的内容为：" + string(lineBytes))
		} else {
			log.Println("没有内容")
		}
		if terr == io.EOF {
			break
		}
	}

}

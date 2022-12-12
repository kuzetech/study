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
	file.Write([]byte("1111\n"))
	file.Write([]byte("2222\n"))
	file.Write([]byte("3333\n"))
	file.Close()
}

func Test_read(t *testing.T) {
	file, _ := os.Open("./files/read.txt")
	r := bufio.NewReader(file)
	for {
		lineBytes, terr := r.ReadBytes('\n')
		if len(lineBytes) > 0 {
			// 删除输出 '\n'
			log.Println("读取到的内容为：" + string(lineBytes[:(len(lineBytes)-1)]))
		}
		if terr == io.EOF {
			break
		}
	}

}

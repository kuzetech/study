package fileSeek

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func test() {
	var filePath = "./test"

	file, _ := os.Create(filePath)
	defer file.Close()

	var byte1 = []byte("1111\n")
	file.Write(byte1)

	var byte2 = []byte("2222\n")
	file.Write(byte2)

	var byte3 = []byte("3333\n")
	file.Write(byte3)

	//state, _ := os.Stat(filePath)
	//var fileSize = state.Size()

	_, err := file.Seek(4, 0)
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)

	for {
		lineBytes, terr := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if terr == io.EOF {
			break
		}
		if len(line) > 0 {
			fmt.Println(line)
		}

	}
}

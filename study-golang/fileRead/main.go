package fileRead

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func test() {
	file, _ := os.Open("./test.log")
	r := bufio.NewReader(file)
	var readByteTotal uint64 = 250
	var readByteCount uint64 = 0
	for {
		lineBytes, terr := r.ReadBytes('\n')
		log.Info().Msg("读取到的内容为：" + string(lineBytes))
		readByteCount = readByteCount + uint64(len(lineBytes)) + 1
		if readByteCount >= readByteTotal {
			log.Info()
			break
		}
		if terr == io.EOF {
			break
		}
	}

}

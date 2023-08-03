package file

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func getOffset(path string) (uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	if fileInfo.Size() == 0 {
		return 0, nil
	}

	for i := int64(1); i <= fileInfo.Size(); i++ {
		offset, err := file.Seek(-1*i, 2)
		if err != nil {
			return 0, err
		}

		b := make([]byte, 1)
		_, err = file.Read(b)
		if err != nil {
			return 0, err
		}

		if string(b) == "\n" {
			return uint64(offset + 1), nil
		}
	}

	return 0, nil
}

func Test_read_last_line(t *testing.T) {
	assertions := require.New(t)

	result, err := getOffset("/Users/kuze/code/study/study-golang/file/1.log")
	assertions.Nil(err)

	fmt.Println(result)

}

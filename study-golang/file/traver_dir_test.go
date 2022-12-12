package file

import (
	"io/ioutil"
	"testing"
)

func Test_traver_dir(t *testing.T) {
	dir, _ := ioutil.ReadDir("./files")
	for _, fi := range dir {
		if fi.IsDir() {
			println("文件夹：" + fi.Name())
		} else {
			println("文件：" + fi.Name())
		}
	}
}

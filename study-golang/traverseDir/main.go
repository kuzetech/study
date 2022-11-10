package traverseDir

import "io/ioutil"

func traverseDir(path string) {
	dir, _ := ioutil.ReadDir(path)
	for _, fi := range dir {
		if fi.IsDir() {
			println("文件夹：" + fi.Name())
		} else {
			println("文件：" + fi.Name())
		}
	}
}

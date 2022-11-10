package filePath

import "path/filepath"

func test() {
	path := "./*/*.log"

	abs, _ := filepath.Abs(path)

	println(abs)
}

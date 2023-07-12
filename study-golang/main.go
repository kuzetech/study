package main

import (
	"log"
	"techkuze.com/bigdata/study/study-golang/file"
)

func main() {
	path, err := file.FindInodePath(493098715)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(path)
}

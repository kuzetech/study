package main

import (
	"log"
)

func main() {

	var m = make(map[interface{}]interface{})

	m[1] = 1
	m["a"] = "a"

	log.Println("test", m)
}

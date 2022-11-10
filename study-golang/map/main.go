package _map

import "log"

func test2(a map[int]int) {
	a[1] = 1
}

func test() {
	test := make(map[int]int)
	test[0] = 0
	test[1] = 0
	test[2] = 0
	log.Println(test)
	test2(test)
	log.Println(test)

}

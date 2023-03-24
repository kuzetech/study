package _map

import (
	"log"
	"testing"
)

// 实参
func deleteItem1(m map[string]interface{}, key string) {
	delete(m, key)
}

// 实参
func deleteItem2(m *map[string]interface{}, key string) {
	delete(*m, key)
}

// 实参
func deleteItem3(m interface{}, key string) {
	m2 := m.(map[string]interface{})
	delete(m2, key)
}

// 实参
func deleteItem4(m *interface{}, key string) {
	m2 := (*m).(map[string]interface{})
	delete(m2, key)
}

func TestMap(t *testing.T) {
	a := make(map[string]interface{})
	a["1"] = 1
	a["2"] = 2
	a["3"] = 3

	var b interface{} = a
	deleteItem4(&b, "3")

	log.Println(a)
}

package sortslice

import (
	"container/list"
	"fmt"
	"sort"
)

var strSlice []int
var strSlice2 []string

func init() {
	strSlice = make([]int, 5)
	strSlice[0] = 1
	strSlice[1] = 5
	strSlice[2] = 6
	strSlice[3] = 2
	strSlice[4] = 14

	strSlice2 = make([]string, 8)
	strSlice2[0] = "SuccessWal-99999"
	strSlice2[1] = "SuccessWal-99996"
	strSlice2[2] = "SuccessWal-99993"
	strSlice2[3] = "SuccessWal-99997"

	strSlice2[4] = "SuccessWal-00005"
	strSlice2[5] = "SuccessWal-00003"
	strSlice2[6] = "SuccessWal-00004"
	strSlice2[7] = "SuccessWal-00016"
}

func getBatchIdStrFromSuccessWalKey(key string) string {
	batchIdStr := key[len("SuccessWal-"):]
	return batchIdStr
}

func sortList() {
	//fmt.Printf("strSlice: %v\n", strSlice)
	//sort.Ints(strSlice)
	//fmt.Printf("strSlice: %v\n", strSlice)

	fmt.Printf("strSlice: %v\n", strSlice2)
	sort.Strings(strSlice2)
	fmt.Printf("strSlice: %v\n", strSlice2)

	keyList := list.New()
	for _, s := range strSlice2 {
		keyList.PushBack(s)
	}

	first := getBatchIdStrFromSuccessWalKey(strSlice2[0])[0:1]
	last := getBatchIdStrFromSuccessWalKey(strSlice2[len(strSlice2)-1])[0:1]
	if first == "0" && last == "9" {
		for i := len(strSlice2) - 1; getBatchIdStrFromSuccessWalKey(strSlice2[i])[0:1] == "9"; i = i - 1 {
			keyList.PushFront(keyList.Back().Value.(string))
			keyList.Remove(keyList.Back())
		}
	}

	for i := keyList.Front(); i != nil; i = i.Next() {
		s := i.Value.(string)
		println(s)
	}

}

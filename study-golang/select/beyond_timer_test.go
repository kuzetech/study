package _select

import (
	"log"
	"testing"
	"time"
)

func Test_beyond_timer(t *testing.T) {

	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	for true {
		select {
		case currentTime := <-ticker.C:
			log.Println(currentTime)
			time.Sleep(time.Second * 5)
		}
	}

}

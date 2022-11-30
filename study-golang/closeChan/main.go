package closeChan

import (
	"log"
	"time"
)

func testCloseChan() {

	testChan := make(chan struct{})

	go func() {
		log.Println("启动监听线程")
		for {
			select {
			case <-testChan:
				log.Println("chan 接收到消息")
				return
			}
		}
	}()

	time.Sleep(time.Second * 2)

	log.Println("准备关闭 chan")
	close(testChan)

	time.Sleep(time.Second * 2)
	log.Println("主线程结束")
}

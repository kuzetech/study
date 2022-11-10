package ticker

import (
	"fmt"
	"time"
)

func test() {
	// Ticker 包含一个通道字段C，每隔时间段 d 就向该通道发送当时系统时间。
	// 它会调整时间间隔或者丢弃 tick 信息以适应反应慢的接收者。
	// 如果d <= 0会触发panic。关闭该 Ticker 可以释放相关资源。
	ticker1 := time.NewTicker(5 * time.Second)

	// 一定要调用Stop()，回收资源
	defer ticker1.Stop()

	go func(t *time.Ticker) {
		for {
			// 每5秒中从chan t.C 中读取一次
			currentTime := <-t.C
			fmt.Println("Ticker:", currentTime.Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	time.Sleep(30 * time.Second)
	fmt.Println("ok")
}

package timer

import (
	"fmt"
	"time"
)

func runTask(flushTimer *time.Timer) {
	for {
		t := <-flushTimer.C
		fmt.Println("触发事件，时间为：", t)
		flushTimer.Reset(10 * time.Second)
	}

}

func test() {
	// 仅执行一次
	flushTimer := time.NewTimer(10 * time.Second)

	go runTask(flushTimer)

	time.Sleep(15 * time.Second)
	fmt.Println("主线程醒过来了，重置了定时器，时间节点为：", time.Now())
	flushTimer.Reset(10 * time.Second)

	time.Sleep(100 * time.Second)

}

func test2() {
	// 仅执行一次
	tChannel := time.After(3 * time.Second) // 其内部其实是生成了一个Timer对象
	select {
	case <-tChannel:
		fmt.Println("3秒执行任务")
	}
}

func test3() {
	timer1 := time.NewTimer(5 * time.Second)
	timer1.Stop()

	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05"))
	go func(t *time.Timer) {
		times := 0
		for {
			<-t.C
			fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))

			// 从t.C中获取数据，此时time.Timer定时器结束。如果想再次调用定时器，只能通过调用 Reset() 函数来执行
			// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。
			// 如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			times++
			// 调用 reset 重发数据到chan C
			fmt.Println("调用 reset 重新设置一次timer定时器，并将时间修改为2秒")
			t.Reset(2 * time.Second)
			if times > 3 {
				fmt.Println("调用 stop 停止定时器")
				t.Stop()
			}
		}
	}(timer1)

	time.Sleep(30 * time.Second)
	fmt.Println("结束时间：", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("ok")
}

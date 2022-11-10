package cond

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func test() {
	var c *sync.Cond = sync.NewCond(&sync.Mutex{})

	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 当主线程持有锁，这里还能继续获取到锁？
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已经准备就绪\n", i)
			// 广播唤醒裁判
			c.Broadcast()
		}(i)
	}

	// 调用 wait 方法时，必须要持有锁
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	log.Println("裁判：所有运动员准备完毕，比赛即将开始")

}

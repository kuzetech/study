package once

import (
	"net"
	"sync"
	"time"
)

var mu sync.Mutex
var conn net.Conn

func getConnByMutex() net.Conn {
	mu.Lock()
	defer mu.Unlock()

	if conn != nil {
		return conn
	}
	newConn, _ := net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return newConn
}

var once sync.Once

func getConnByOnce() {
	once.Do(func() {
		conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	})
}

func main() {

	// 这个方式有性能问题，即使创建好了也需要抢锁
	conn = getConnByMutex()

	// 即使多次调用也只有第一次能执行成功，还没有抢锁的损耗
	getConnByOnce()

}

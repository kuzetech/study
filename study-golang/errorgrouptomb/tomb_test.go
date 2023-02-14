package errorgrouptomb

import (
	"gopkg.in/tomb.v2"
	"testing"
	"time"
)

func TestTomb(t *testing.T) {
	var to tomb.Tomb

	to.Go(func() error {
		for {
			select {
			case <-time.After(time.Second * 2):
				println(1)
			case <-to.Dying():
				println(2)
				return nil
			}
		}
	})

	to.Go(func() error {
		time.Sleep(time.Second * 6)
		println("down1")
		//return errors.New("test")
		return nil
	})

	// 等待所有的协程结束
	// 需要使用 tomb.Dying() 获知其他协程出错
	to.Wait()

	println("down2")
}

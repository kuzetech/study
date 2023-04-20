package errorgroup

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestErrorGroup(t *testing.T) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	eg, ctx := errgroup.WithContext(ctx)
	defer cancel()

	eg.Go(func() error {
		for {
			select {
			case <-time.After(time.Second * 2):
				println(1)
			case <-ctx.Done():
				println("down long")
				return errors.New(ctx.Err().Error())
			}
		}
	})

	eg.Go(func() error {
		time.Sleep(time.Second * 5)
		println("down short")
		return errors.New("test1")
	})

	// 等待所有的协程结束
	// 仅当所有协程都结束了才返回 nil
	// 如果其中一个协程出现错误，也需要等待其他协程结束了，才会返回第一个协程出现的错误，意思是即使多个协程出错，也只返回第一个错误
	// 其他的线程可以监听 Context 获知 group 结束

	err := eg.Wait()
	if err != nil {
		fmt.Println("err =", err)
	}
	println("down main")

}

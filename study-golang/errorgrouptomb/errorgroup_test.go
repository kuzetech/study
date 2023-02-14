package errorgrouptomb

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestErrorGroup(t *testing.T) {
	eg := errgroup.Group{}

	eg.Go(func() error {
		fmt.Println("go1")
		return nil
	})

	eg.Go(func() error {
		fmt.Println("go2")
		err := errors.New("go2 err")
		return err
	})

	// 等待所有的协程结束，并返回错误信息
	// 特别注意，其中一个协程错误，其他协程并不会停止
	err := eg.Wait()
	if err != nil {
		fmt.Println("err =", err)
	}
}

package test

import "testing"

func TestReverse(t *testing.T) {
	str := "abc"
	revStr1 := Reverse(str)
	revStr2 := Reverse(revStr1)
	if str != revStr2 {
		// error 方法报错后, 会继续向下执行，需要格式化信息可以调用 t.Errorf()
		t.Error("error")

		// fatal 方法报错后, 会退出测试，需要格式化信息可以调用 t.Fatalf()
		// t.Fatal("fatal")

		// 测试中断, 但是测试结果不会十遍
		// t.Skip("skip")

		// 输出调试信息，需要格式化信息可以调用 t.Logf()
		// t.Log("log")
	}
	// 可启动多个子测试, 子测试之间并行运行
	for _, str = range []string{"abcd", "aceb"} {
		// 第一个参数为子测试的标识
		t.Run(str, func(t *testing.T) {
			revStr1 := Reverse(str)
			revStr2 := Reverse(revStr1)
			if str != revStr2 {
				t.Error("error")
			}
		})
	}
}

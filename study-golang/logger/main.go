package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Student struct {
	name string
	age  int
}

func test() {

	// 全局设置
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// 局部设置
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	student := Student{name: "a", age: 1}

	// {"level":"info","test":{},"time":"2022-10-27T09:58:28+08:00","caller":"/Users/huangsw/code/study/study-golang/logger/main_linux.go:23","message":"--------"}
	logger.Info().Interface("test", student).Msg("--------")
}

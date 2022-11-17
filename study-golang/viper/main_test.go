package viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"testing"
)
import "github.com/spf13/viper"

func Test_main(t *testing.T) {
	v := viper.New()

	v.SetConfigName("test")
	v.SetConfigType("yaml")

	// 在目录下查找 test.yaml 的配置文件
	// 多个路径的情况下，读取找到的第一个文件
	v.AddConfigPath("./a")
	v.AddConfigPath("./b")

	err := v.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	filed := v.GetString("test")
	fmt.Println(filed)

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(v.GetString("test"))
	})
	v.WatchConfig()

	select {}

}

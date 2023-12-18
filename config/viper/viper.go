/**
 * @Description: 获取配置文件
 * @Author Lee
 * @Date 2023/12/7 10:01
 **/

package viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init(configPath string) {
	//制定配置文件的路径
	viper.SetConfigFile(configPath)
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		panic(fmt.Errorf("配置文件读取失败，错误信息:%v", err))
	}
	//监听修改
	viper.WatchConfig()
	//为配置修改增加一个回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
}

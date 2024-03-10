package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	SensitiveWords []string

	MessageQueueLen = 1024
)

func initConfig() {
	viper.SetConfigName("chatroom")
	viper.AddConfigPath(RootDir + "/config")
	//尝试读取config中的内容
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	//用yaml中的内容初始化全局变量
	SensitiveWords = viper.GetStringSlice("sensitive")
	MessageQueueLen = viper.GetInt("message-queue")
	//监听Config
	viper.WatchConfig()
	//一旦Config改变则更新敏感词
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.ReadInConfig()
		SensitiveWords = viper.GetStringSlice("sensitive")
	})
}

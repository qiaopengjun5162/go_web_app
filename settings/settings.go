package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	// 设置默认值
	viper.SetDefault("fileDir", "./")
	// 读取配置文件
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")          // SetConfigType设置远端源返回的配置类型，例如:“json”。
	viper.AddConfigPath(".")             // 还可以在工作目录中查找配置

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig failed, error: %v\n", err)
		return
	}

	// 实时监控配置文件的变化 WatchConfig 开始监视配置文件的更改。
	viper.WatchConfig()
	// OnConfigChange设置配置文件更改时调用的事件处理程序。
	// 当配置文件变化之后调用的一个回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return
}

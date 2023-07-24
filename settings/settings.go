package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("conf/config.yaml") // 指定配置文件
	//viper.SetConfigName("config") // 指定配置文件名称(不需要带后缀)
	//viper.SetConfigType("yaml")   // 指定配置文件类型(专用于从远程获取配置信息时指定配置文件类型)
	viper.AddConfigPath(".")   // 指定查找配置文件的路径(使用相对路径)
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
	return
}

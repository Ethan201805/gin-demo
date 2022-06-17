package core

import (
	"fmt"
	"gin-web-demo/Test/gin-demo/admin/core/internal"
	"gin-web-demo/Test/gin-demo/admin/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper() *viper.Viper {
	config := internal.ConfigDefaultFile
	if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
		switch gin.Mode() {
		case gin.DebugMode:
			config = internal.ConfigDefaultFile
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径1为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
		case gin.ReleaseMode:
			config = internal.ConfigReleaseFile
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径2为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
		case gin.TestMode:
			config = internal.ConfigTestFile
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径3为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
		}
	} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
		config = configEnv
		fmt.Printf("您正在使用%s环境变量,config的路径4为%s\n", internal.ConfigEnv, config)
	}
	v := viper.New()
	v.AddConfigPath("/Users/wangxiang/Ethan/go/gin-web-demo/Test/gin-demo/admin")
	v.SetConfigFile(config)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.ET_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.ET_CONFIG); err != nil {
		fmt.Println(err)
	}
	//global.ET_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

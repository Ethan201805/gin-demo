package main

import (
	"gin-web-demo/Test/gin-demo/admin/core"
	"gin-web-demo/Test/gin-demo/admin/global"
	"gin-web-demo/Test/gin-demo/admin/initialize"
	"gin-web-demo/Test/gin-demo/admin/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//通过viper读取配置文件
	global.ET_VP = core.Viper()
	//初始化数据库
	global.ET_DB = initialize.Gorm() // gorm连接数据库
	//初始化日志库

	//初始化路由
	router := router.InitRouter()
	//Listen and serve on 8090
	router.Run(":8090")
}

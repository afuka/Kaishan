package main

import (
	"github.com/json-iterator/go/extra"
	"kaishan/core/handlers/conf"
	"kaishan/core/handlers/httpser"
	"kaishan/core/handlers/ice"
	"kaishan/core/handlers/log"
	"kaishan/core/handlers/mysql"
	"kaishan/core/handlers/redis"
)

// initHandler 初始化
func initHandler() {
	log.InitLogger("output/logs") // 初始化日志
	extra.RegisterFuzzyDecoders() // 滴滴开源的第三方json编码库
	conf.InitConfig(confPath) // 配置文件初始化
	redis.InitRedis() // 初始化redis
	mysql.InitMysql() // 初始化mysql
	ice.InitIce() // 初始化id生成器
	httpser.InitHttpSer() // http服务端
}

// closeHandler 关闭具柄
func closeHandler()  {
	httpser.Close()
	redis.Close()
	mysql.Close()
}
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
	extra.RegisterFuzzyDecoders() // 滴滴开源的第三方json编码库
	conf.InitConfig(confPath) // 配置文件初始化
	redis.InitRedis() // 初始化redis
	mysql.InitMysql() // 初始化mysql
	ice.InitIce() // 初始化id生成器
	log.InitLogger() // 初始化日志
	httpser.InitHttpSer(quit) // http服务端
}

// closeHandler 关闭所有具柄
func closeHandler()  {
	redis.Close()
	mysql.Close()

	quit <- true // 关闭 http ser
}
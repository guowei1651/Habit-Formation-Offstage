package database

import (
	"log"
	"hf/config"
)

func OpenConnectPool () {
	log.Printf("开始创建数据库连接")

	log.Printf("数据库配置参数:%v", config)
	sqlOpen()

	log.Printf("创建数据库连接结束")
}
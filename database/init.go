package database

import (
	"log"
	hfConfig "hf/config"

	"database/sql"
)

type DBProcess interface {
	ConnectPing()
	ConnectOpen()
	GetPool() (db *sql.DB)
}

var DBConnectPool *sql.DB

func OpenConnectPool () {
	log.Printf("开始创建数据库连接")
	log.Printf("数据库配置参数:%v", hfConfig.Config)
	var pool DBProcess
	if (hfConfig.Config.DBConfig.Schmea == "postgres") {
		pool = &Postgres{}
	} else if (hfConfig.Config.DBConfig.Schmea == "mysql") {
		pool = &Mysql{}
	} else {
		panic("配置schema配置不支持")
	}

	log.Printf("OpenConnectPool ConnectOpen pool: %v", pool)
	pool.ConnectOpen()

	log.Printf("OpenConnectPool ConnectPing pool: %v", pool)
	pool.ConnectPing()

	log.Printf("OpenConnectPool GetPool pool: %v", pool)
	DBConnectPool = pool.GetPool()

	log.Printf("创建数据库连接结束")
}
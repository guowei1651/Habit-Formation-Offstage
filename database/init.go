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
	var pool *DBProcess
	if (hfConfig.Config.DBConfig.Schmea == "postgres") {
		pool = &Postgres{}
	} else if (hfConfig.Config.DBConfig.Schmea == "mysql") {
		pool = &Mysql{}
	} else {
		panic("配置schema配置不支持")
	}
	
	pool.ConnectOpen()

	pool.ConnectPing()

	DBConnectPool = pool.GetPool()

	log.Printf("创建数据库连接结束")
}
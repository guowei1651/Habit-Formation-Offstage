package database

import (
	"log"
	hfConfig "hf/config"

	"database/sql"
)

type DB interface {
	connectPing()
	connectOpen()
}

var DBConnectPool *sql.DB

func OpenConnectPool () {
	log.Printf("开始创建数据库连接")
	log.Printf("数据库配置参数:%v", hfConfig.Config)
	var pool *DB
	if (hfConfig.Config.DBConfig.schmea == "postgres") {
		pool = &Postgres{}
	}
	if (hfConfig.Config.DBConfig.schmea == "mysql") {
		pool = &Mysql{}
	}
	
	pool.connectOpen()

	pool.connectPing()

	log.Printf("创建数据库连接结束")
}
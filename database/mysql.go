package database

import (
    "log"

	"database/sql"
)

type Mysql struct {
    DBConnectPool *sql.DB
}

func (m Mysql) connectPing() {
	log.Printf("未实现")
}

func (m Mysql) connectOpen() {
	log.Printf("未实现")
}
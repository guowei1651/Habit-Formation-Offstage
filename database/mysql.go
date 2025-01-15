package database

import (
    "log"

	"database/sql"
)

type Mysql struct {
    Pool *sql.DB
}

func (m Mysql) ConnectPing() {
	log.Printf("未实现")
}

func (m Mysql) ConnectOpen() {
	log.Printf("未实现")
}

func (m Mysql) GetPool() (Pool *sql.DB) {
	log.Printf("未实现")
	panic("未实现")
}
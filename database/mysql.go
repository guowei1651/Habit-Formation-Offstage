package database

import (
    "log"

	"database/sql"
)

type Mysql struct {
    db *sql.DB
}

func (m Mysql) connectPing() {
	log.Printf("未实现")
}

func (m Mysql) connectOpen() {
	log.Printf("未实现")
}

func (m Mysql) getPool () (*sql.DB){
	log.Printf("未实现")
	panic("未实现")
}
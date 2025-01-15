// https://jianghushinian.cn/2023/06/05/how-to-use-database-sql-to-operate-database-in-go/

package database

import (
    "log"
    "time"
    "context"
    "database/sql"
	_ "github.com/lib/pq"
)

type Postgres struct {
    db *sql.DB
}

func (p Postgres) connectPing() {
    ctx, stop := context.WithCancel(context.Background())
    defer stop()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	log.Printf("db ping db-> ", p.db)
	if err := p.db.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
        panic(err)
	}
}

func (p Postgres) connectOpen() {
    log.Printf("开始打开Postgres链接")
    var err error
    p.db, err = sql.Open("postgres", "host=172.25.1.22 port=5432 user=appsmith password=appsmith dbname=appsmith sslmode=disable")
    // db, err = sql.Open("postgres", "postgres://appsmith:appsmith@172.25.1.22:5432/appsmith?sslmode=disable")
    log.Printf("open db complete -> ", p.db)
    p.db.SetConnMaxIdleTime(30*1000)
    p.db.SetConnMaxLifetime(10*1000)
    p.db.SetMaxIdleConns(10)
    p.db.SetMaxOpenConns(20)
    if err != nil {
        log.Fatal("open db connect fail -> ", err)
        panic(err)
    }
    log.Printf("打开Postgres链接完成")
}